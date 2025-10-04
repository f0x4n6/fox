//go:build !minimal

package ui

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"

	_ "github.com/gdamore/tcell/v2/encoding"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ai/chat"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/opt/ui/widgets"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/mode"
	"github.com/cuhsat/fox/internal/pkg/user/bag"
	"github.com/cuhsat/fox/internal/pkg/user/history"
	"github.com/cuhsat/fox/internal/pkg/user/plugins"
)

const (
	delta = 1 // lines
	wheel = 5 // lines
)

const (
	bpPrefix = "ESC[200~" // bracketed paste start
	bpSuffix = "ESC[201~" // bracketed paste end
)

type UI struct {
	root tcell.Screen

	state   *opt.State
	chats   *sync.Map
	themes  *themes.Themes
	plugins *plugins.Plugins

	title   *widgets.Title
	view    *widgets.View
	prompt  *widgets.Prompt
	overlay *widgets.Overlay
}

func Start(args []string, invoke types.Invoke) {
	hs := heapset.New(args)
	defer hs.ThrowAway()

	hi := history.New()
	defer hi.Close()

	bg := bag.New()
	defer bg.Close()

	ui := create()
	defer ui.delete()

	ui.run(hs, hi, bg, invoke)
}

func create() *UI {
	runewidth.CreateLUT()

	root, err := tcell.NewScreen()

	if err != nil {
		log.Panicln(err)
	}

	err = root.Init()

	if err != nil {
		log.Panicln(err)
	}

	if !flags.Get().Opt.NoMouse {
		root.EnableMouse(tcell.MouseDragEvents)
	}

	root.EnablePaste()

	state := opt.NewState(root)

	ui := UI{
		root: root,

		state:   state,
		chats:   new(sync.Map),
		themes:  themes.New(state.Theme()),
		plugins: plugins.New(),

		title:   widgets.NewTitle(state),
		view:    widgets.NewView(state),
		prompt:  widgets.NewPrompt(state),
		overlay: widgets.NewOverlay(state),
	}

	root.SetCursorStyle(widgets.Cursor, themes.Cursor)
	root.SetStyle(themes.Terminal)
	root.Sync()

	ui.render(nil)
	ui.changeMode(flags.Get().UI.Mode)

	return &ui
}

func (ui *UI) delete() {
	ui.chats.Range(func(_, c any) bool {
		c.(*chat.Chat).Close()
		return true
	})

	if ui.plugins != nil {
		plugins.Close()
	}

	ui.overlay.Close()
	ui.root.Fini()
	ui.state.Save()
}

func (ui *UI) run(hs *heapset.HeapSet, hi *history.History, bg *bag.Bag, util types.Invoke) {
	hs.SetCallback(func() {
		_ = ui.root.PostEvent(tcell.NewEventInterrupt(ui.state.IsFollow()))
	})

	events := make(chan tcell.Event, 128)
	closed := make(chan struct{})

	go ui.root.ChannelEvents(events, closed)
	go ui.overlay.Listen()

	ui.invoke(hs, util)

	esc := false

	for {
		select {
		case _ = <-closed:
			return // channels closed

		case ev := <-events:
			if ev == nil {
				return // terminal closed
			}

			w, h := ui.root.Size()

			_, heap := hs.Heap()

			switch ev := ev.(type) {
			case *tcell.EventInterrupt:
				if v, ok := ev.Data().(bool); ok && v {
					ui.view.ScrollLast()
				}

			case *tcell.EventClipboard:
				if ui.state.Mode().Static() {
					continue
				}

				v := string(ev.Data())
				v = strings.TrimPrefix(v, bpPrefix)
				v = strings.TrimSuffix(v, bpSuffix)
				ui.prompt.SetValue(v)

			case *tcell.EventResize:
				ui.root.Sync()
				ui.view.Reset()

			case *tcell.EventError:
				ui.overlay.SendError(ev.Error())

			case *tcell.EventMouse:
				ui.handleMouse(hs, heap, ev)

			case *tcell.EventKey:
				mods := ev.Modifiers()

				pageW := w - 1 // minus text abbreviation
				pageH := h - 2 // minus title and status

				if ui.state.IsNavi() {
					pageW -= text.Dec(heap.Count()) + 1
				}

				if ev.Key() != tcell.KeyEscape {
					esc = false // reset
				}

				switch ev.Key() {
				case tcell.KeyEscape:
					if esc {
						return // twice to exit
					} else if ui.state.Mode().Prompt() {
						ui.changeBack()
					}

					esc = true

				case tcell.KeyTab:
					if mods&tcell.ModShift == 0 {
						ui.nextTab(hs, heap)
					} else {
						ui.prevTab(hs, heap)
					}

				case tcell.KeyF1:
					ui.view.Reset()
					hs.OpenHelp()

				case tcell.KeyF2:
					hs.Counts()
					ui.changeMode(mode.Default)

				case tcell.KeyF3:
					hs.Entropy(0.0, 1.0)
					ui.changeMode(mode.Default)

				case tcell.KeyF4:
					hs.Strings(3, math.MaxInt, true, nil)
					ui.changeMode(mode.Default)

				case tcell.KeyF5:
					hs.HashSum(types.MD5)
					ui.changeMode(mode.Default)

				case tcell.KeyF6:
					hs.HashSum(types.SHA1)
					ui.changeMode(mode.Default)

				case tcell.KeyF7:
					hs.HashSum(types.SHA256)
					ui.changeMode(mode.Default)

				case tcell.KeyF8:
					ui.runAssistant(hs, heap)

				case tcell.KeyF9, tcell.KeyF10, tcell.KeyF11, tcell.KeyF12:
					fallthrough
				case tcell.KeyF13, tcell.KeyF14, tcell.KeyF15, tcell.KeyF16:
					fallthrough
				case tcell.KeyF17, tcell.KeyF18, tcell.KeyF19, tcell.KeyF20:
					fallthrough
				case tcell.KeyF21, tcell.KeyF22, tcell.KeyF23, tcell.KeyF24:
					ui.runPlugin(hs, heap, strings.ToLower(ev.Name()))

				case tcell.KeyUp:
					switch {
					case ui.state.Mode().Prompt():
						ui.prompt.SetValue(hi.PrevLine())
					case mods&tcell.ModShift != 0 && mods&tcell.ModCtrl != 0:
						ui.view.ScrollStart()
					case mods&tcell.ModShift != 0:
						ui.view.ScrollUp(pageH)
					default:
						ui.view.ScrollUp(delta)
					}

				case tcell.KeyDown:
					switch {
					case ui.state.Mode().Prompt():
						ui.prompt.SetValue(hi.NextLine())
					case mods&tcell.ModShift != 0 && mods&tcell.ModCtrl != 0:
						ui.view.ScrollEnd()
					case mods&tcell.ModShift != 0:
						ui.view.ScrollDown(pageH)
					default:
						ui.view.ScrollDown(delta)
					}

				case tcell.KeyLeft:
					switch {
					case ui.state.Mode().Prompt():
						if mods&tcell.ModCtrl != 0 {
							ui.prompt.MoveStart()
						} else {
							ui.prompt.Move(-1)
						}
					case mods&tcell.ModCtrl != 0:
						ui.prevTab(hs, heap)
					case mods&tcell.ModShift != 0:
						ui.view.ScrollLeft(pageW)
					default:
						ui.view.ScrollLeft(delta)
					}

				case tcell.KeyRight:
					switch {
					case ui.state.Mode().Prompt():
						if mods&tcell.ModCtrl != 0 {
							ui.prompt.MoveEnd()
						} else {
							ui.prompt.Move(+1)
						}
					case mods&tcell.ModCtrl != 0:
						ui.nextTab(hs, heap)
					case mods&tcell.ModShift != 0:
						ui.view.ScrollRight(pageW)
					default:
						ui.view.ScrollRight(delta)
					}

				case tcell.KeyHome:
					ui.view.ScrollStart()

				case tcell.KeyPgUp:
					ui.view.ScrollUp(pageH)

				case tcell.KeyPgDn:
					ui.view.ScrollDown(pageH)

				case tcell.KeyEnd:
					ui.view.ScrollEnd()

				case tcell.KeyCtrlCarat:
					ui.changeTheme()

				case tcell.KeyCtrlSpace:
					ui.changeMode(mode.Goto)

				case tcell.KeyCtrlO:
					ui.changeMode(mode.Open)

				case tcell.KeyCtrlG:
					ui.changeMode(mode.Goto)

				case tcell.KeyCtrlL:
					ui.changeMode(mode.Less)

				case tcell.KeyCtrlF:
					ui.changeMode(mode.Grep)

				case tcell.KeyCtrlX:
					ui.changeMode(mode.Hex)

				case tcell.KeyCtrlP:
					ui.state.TogglePinned()

				case tcell.KeyCtrlT:
					ui.state.ToggleFollow()

				case tcell.KeyCtrlN:
					ui.state.ToggleNavi()
					ui.view.Preserve()

				case tcell.KeyCtrlW:
					ui.state.ToggleWrap()
					ui.view.Preserve()

				case tcell.KeyCtrlJ:
					if heap.ModContext(-1) {
						ui.view.Reset()
					}

				case tcell.KeyCtrlK:
					if heap.ModContext(+1) {
						ui.view.Reset()
					}

				case tcell.KeyCtrlV:
					ui.root.GetClipboard()

				case tcell.KeyCtrlA:
					if !ui.state.Mode().Static() && hs.Merge() {
						ui.overlay.SendInfo("Merged all open files")
					}

				case tcell.KeyCtrlU:
					//if !ui.state.Mode().Static() && hs.Merge() {
					//	ui.overlay.SendInfo("Created super timeline")
					//}

				case tcell.KeyCtrlC:
					if !ui.state.Mode().Static() {
						ui.root.SetClipboard(heap.Bytes())
						ui.overlay.SendInfo(fmt.Sprintf("%s copied to clipboard", heap.String()))
					}

				case tcell.KeyCtrlS:
					if !ui.state.Mode().Static() && bg.Put(heap) {
						ui.overlay.SendInfo(fmt.Sprintf("%s saved to %s", heap.String(), bg.String()))
					}

				case tcell.KeyCtrlB:
					if fs.Exists(bg.Path) {
						ui.view.Reset()
						hs.OpenFile(bg.Path, bg.Path, bg.Path, types.Ignore)
					} else {
						ui.overlay.SendError(fmt.Sprintf("%s not found", bg.Path))
					}

				case tcell.KeyCtrlQ:
					ui.view.Reset()

					if hs.CloseHeap() == nil {
						return // exit
					}

				case tcell.KeyCtrlZ:
					ui.runShell()

				case tcell.KeyEnter:
					l := ui.prompt.ReadLine()
					m := ui.state.Mode()

					if m.Prompt() && len(strings.TrimSpace(l)) == 0 {
						continue
					}

					hi.AddLine(l)

					switch m {
					case mode.Less, mode.Hex:
						ui.view.ScrollLine()

					case mode.Grep:
						ui.view.Reset()
						heap.AddFilter(l, 0, 0)
						ui.changeMode(mode.Less)

					case mode.Goto:
						ui.view.Goto(l)
						ui.changeMode(ui.state.Last())

					case mode.Open:
						ui.state.Call(func() {
							hs.Open(l)
						})
						ui.changeMode(ui.state.Last())

					case mode.Chat:
						ui.view.Reset()
						ui.state.Call(func() {
							if v, ok := ui.chats.Load(heap.String()); ok {
								v.(*chat.Chat).Query(l, true)
							}
						})

					default:
						plugins.Input <- l
						ui.changeMode(ui.state.Last())
					}

				case tcell.KeyDelete:
					ui.prompt.DelRune(widgets.After)

				case tcell.KeyBackspace, tcell.KeyBackspace2:
					switch {
					case len(ui.prompt.GetValue()) > 0:
						ui.prompt.DelRune(widgets.Before)
					case ui.state.Mode().Prompt():
						ui.changeBack()
					case len(heap.Patterns()) > 0 && !ui.state.Mode().Static():
						ui.view.Reset()
						heap.DelFilter()
					}

				default:
					r := ev.Rune()

					switch r {
					case 0: // error
						continue

					case 32: // space
						if ui.prompt.Locked() {
							ui.view.ScrollDown(pageH)
						} else {
							ui.prompt.AddRune(r)
						}

					default: // all other keys
						if ui.state.Mode() == mode.Less {
							ui.changeMode(mode.Grep)
						}

						if unicode.IsPrint(r) {
							ui.prompt.AddRune(r)
						}
					}
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

	case types.Strings:
		hs.Strings(
			flg.Strings.Min,
			flg.Strings.Max,
			flg.Strings.Ioc,
			flg.Strings.Re,
		).CloseOther()

	case types.Hash:
		hs.HashSum(
			flg.Hash.Algos.Value...,
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
