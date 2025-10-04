//go:build !minimal

package ui

import (
	"fmt"
	"log"

	"github.com/cuhsat/fox/internal/opt/ai"
	"github.com/cuhsat/fox/internal/opt/ai/chat"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/opt/ui/widgets"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/mode"
)

type Direction int

const (
	Prev Direction = iota
	Next
)

func format(l, r string) string {
	return text.Title(l, r, !flags.Get().UI.Legacy)
}

func (ui *UI) prevTab(hs *heapset.HeapSet, h *heap.Heap) {
	ui.changeTab(hs, h, Prev)
}

func (ui *UI) nextTab(hs *heapset.HeapSet, h *heap.Heap) {
	ui.changeTab(hs, h, Next)
}

func (ui *UI) changeTab(hs *heapset.HeapSet, h *heap.Heap, d Direction) {
	if hs.Len() <= 1 {
		return
	}

	ui.view.SavePosition(h.Path)

	switch d {
	case Prev:
		h = hs.PrevHeap()
	case Next:
		h = hs.NextHeap()
	}

	ui.view.LoadPosition(h.Path)

	if h.Type == types.Chat {
		ui.changeMode(mode.Chat)
	} else if ui.state.Mode() == mode.Chat {
		ui.changeMode(ui.state.Last())
	}
}

func (ui *UI) changeMode(m mode.Mode) {
	// check for AI support
	if m == mode.Chat && !ai.Check() {
		ui.overlay.SendError("Assistant is not available")
		return
	}

	if !ui.state.SwitchMode(m) {
		return
	}

	// former mode
	if ui.state.Last().Prompt() {
		ui.prompt.SetValue("")
	}

	// actual mode
	ui.prompt.Lock(!m.Prompt())

	// force the cursor off
	if ui.prompt.Locked() {
		ui.state.Root.HideCursor()
	}

	if ui.state.Last().Static() || m.Static() {
		ui.view.Reset()
	}
}

func (ui *UI) changeBack() {
	if ui.state.Last().Prompt() {
		ui.changeMode(mode.Default)
	} else {
		ui.changeMode(ui.state.Last())
	}
}

func (ui *UI) changeTheme() {
	ui.state.ChangeTheme(ui.themes.Cycle())
	ui.root.Fill(' ', themes.Terminal)
	ui.root.Show()
	ui.root.SetCursorStyle(widgets.Cursor, themes.Cursor)
	ui.overlay.SendInfo(fmt.Sprintf("Theme %s", ui.state.Theme()))
}

func (ui *UI) runShell() {
	err := ui.root.Suspend()

	if err != nil {
		sys.Error(err)
		return
	}

	sys.Shell()

	err = ui.root.Resume()

	if err != nil {
		log.Panicln(err)
	}
}

func (ui *UI) runPlugin(hs *heapset.HeapSet, h *heap.Heap, s string) {
	if ui.plugins == nil {
		return // plugins not configured
	}

	p, ok := ui.plugins.Hotkey[s]

	if !ok {
		return // hotkey not configured
	}

	if flags.Get().Opt.NoPlugins {
		ui.overlay.SendError("Plugins deactivated")
		return
	}

	go p.Execute(h.Path, h.Base, func(path, base, dir string) {
		if len(dir) > 0 {
			hs.Open(dir)
		}

		hs.OpenPlugin(path, base, format(base, p.Name))

		ui.state.ForceRender()
		ui.overlay.SendInfo(fmt.Sprintf("%s executed", p.Name))
	})

	if len(p.Mode) > 0 {
		ui.changeMode(mode.Mode(p.Mode))
	}
}

func (ui *UI) runAssistant(hs *heapset.HeapSet, h *heap.Heap) {
	var c *chat.Chat

	if h.Type == types.Chat || h.Type == types.Ignore {
		return
	}

	title := format(h.String(), "assistant")

	if v, ok := ui.chats.Load(title); !ok {
		c = chat.New(ui.state, h)
		ui.chats.Store(title, c)
	} else {
		c = v.(*chat.Chat)
	}

	path := c.File.Name()

	hs.OpenChat(path, path, title)
	ui.changeMode(mode.Chat)
}
