//go:build !minimal

package ui

import (
	"log"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"

	_ "github.com/gdamore/tcell/v2/encoding"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ai/chat"
	"github.com/cuhsat/fox/internal/opt/ui/adapter"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/opt/ui/widgets"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/user/bag"
	"github.com/cuhsat/fox/internal/pkg/user/history"
	"github.com/cuhsat/fox/internal/pkg/user/plugins"
)

const (
	prefix = "ESC[200~" // bracketed paste start
	suffix = "ESC[201~" // bracketed paste end
)

type UI struct {
	root tcell.Screen
	lkey tcell.Key

	state   *opt.State
	chats   *sync.Map
	themes  *themes.Themes
	plugins *plugins.Plugins
	history *history.History
	bag     *bag.Bag

	title   *widgets.Title
	view    *widgets.View
	prompt  *widgets.Prompt
	overlay *widgets.Overlay
}

func Start(args []string, invoke types.Invoke) {
	hs := heapset.New(args)
	defer hs.ThrowAway()

	ui := New(hs, invoke)
	defer ui.Close()

	ui.run(hs)
}

func New(hs *heapset.HeapSet, util types.Invoke) *UI {
	runewidth.CreateLUT()

	root, err := tcell.NewScreen()

	if err != nil {
		log.Panicln(err)
	}

	err = root.Init()

	if err != nil {
		log.Panicln(err)
	}

	root.EnablePaste()

	if !flags.Get().Optional.NoMouse {
		root.EnableMouse(tcell.MouseDragEvents)
	}

	state := opt.NewState(root)

	ui := UI{
		root: root,

		state:   state,
		chats:   new(sync.Map),
		themes:  themes.New(state.Theme()),
		plugins: plugins.New(),
		history: history.New(),
		bag:     bag.New(),

		title:   widgets.NewTitle(state),
		view:    widgets.NewView(state),
		prompt:  widgets.NewPrompt(state),
		overlay: widgets.NewOverlay(state),
	}

	root.SetCursorStyle(widgets.Cursor, themes.Cursor)
	root.SetStyle(themes.Terminal)
	root.Sync()

	ui.invoke(hs, util)
	ui.render(nil)
	ui.changeMode(flags.Get().UI.Mode)

	return &ui
}

func (ui *UI) Close() {
	ui.chats.Range(func(_, c any) bool {
		c.(*chat.Chat).Close()
		return true
	})

	if ui.plugins != nil {
		plugins.Close()
	}

	ui.root.Fini()
	ui.state.Save()

	ui.bag.Close()
	ui.history.Close()
}

func (ui *UI) run(hs *heapset.HeapSet) {
	hs.SetCallback(func() {
		_ = ui.root.PostEvent(tcell.NewEventInterrupt(ui.state.IsFollow()))
	})

	ui.view.Init(adapter.NewFileSystem(ui.state, hs.Open))

	events := make(chan tcell.Event, 128)
	closed := make(chan struct{})

	go ui.root.ChannelEvents(events, closed)
	go ui.overlay.Listen()

	for {
		select {
		case _ = <-closed:
			return // channels closed

		case ev := <-events:
			if ev == nil {
				return // terminal closed
			}

			_, heap := hs.Heap()

			switch ev := ev.(type) {
			case *tcell.EventInterrupt:
				if v, ok := ev.Data().(bool); ok && v {
					ui.view.ScrollLast()
				}

			case *tcell.EventClipboard:
				if !ui.state.Mode().IsStatic() {
					v := string(ev.Data())
					v = strings.TrimPrefix(v, prefix)
					v = strings.TrimSuffix(v, suffix)
					ui.prompt.SetInput(v)
				}

			case *tcell.EventResize:
				ui.root.Sync()
				ui.view.Reset()

			case *tcell.EventFocus:
				if ev.Focused {
					ui.root.Sync()
				}

			case *tcell.EventError:
				ui.overlay.SendError(ev.Error())

			case *tcell.EventMouse:
				ui.handleMouse(hs, heap, ev)

			case *tcell.EventKey:
				if ui.handleKey(hs, heap, ev) {
					return // exit
				}
			}

			ui.render(hs)
		}
	}
}

func (ui *UI) invoke(hs *heapset.HeapSet, util types.Invoke) {
	flg := flags.Get()

	switch util {
	case types.Compare:
		hs.Compare(
			flg.Compare.Git,
		).CloseOther()

	case types.Counts:
		hs.Counts().CloseOther()

	case types.Entropy:
		hs.Entropy(
			flg.Entropy.Min,
			flg.Entropy.Max,
		).CloseOther()

	case types.HashSum:
		hs.HashSum(
			flg.Hash.Algos.Value...,
		).CloseOther()

	case types.Strings:
		hs.Strings(
			flg.Strings.Min,
			flg.Strings.Max,
			flg.Strings.Class,
			flg.Strings.Re,
		).CloseOther()

	case types.Timeline:
		hs.Timeline(
			flg.Timeline.Cef,
		).CloseOther()

	case types.None:
		// normal
	}
}

func (ui *UI) render(hs *heapset.HeapSet) {
	title := fox.Product

	if hs != nil {
		_, heap := hs.Heap()

		if heap.Type == types.Stdin {
			ui.root.Sync() // prevent hiccups
		}

		title = format(title, heap.String())
	}

	ui.root.SetTitle(title)
	ui.root.SetStyle(themes.Terminal)
	ui.root.Clear()

	x, y := 0, 0
	w, h := ui.root.Size()

	for _, base := range [...]widgets.Queueable{
		ui.title,
		ui.view,
		ui.prompt,
	} {
		y += base.Render(hs, x, y, w, h-y)
	}

	ui.overlay.Render(0, 0, w, h)

	ui.root.Show()
}
