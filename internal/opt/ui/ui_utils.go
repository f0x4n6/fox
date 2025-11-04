//go:build !minimal

package ui

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/v3/internal/opt/ai"
	"github.com/cuhsat/fox/v3/internal/opt/ai/chat"
	"github.com/cuhsat/fox/v3/internal/opt/ui/themes"
	"github.com/cuhsat/fox/v3/internal/opt/ui/widgets"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v3/internal/pkg/text"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v3/internal/pkg/types/mode"
)

type Direction int

const (
	Prev Direction = iota
	Next
)

func format(l, r string) string {
	return text.Title(l, r, !flags.Get().UI.Legacy)
}

func (ui *UI) gotoHome() {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveHome())
	default:
		ui.view.ScrollHome()
	}
}

func (ui *UI) gotoEnd() {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveEnd())
	default:
		ui.view.ScrollEnd()
	}
}

func (ui *UI) lineUp() {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveUp(delta))
	default:
		ui.view.ScrollUp(delta)
	}
}

func (ui *UI) lineDown() {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveDown(delta))
	default:
		ui.view.ScrollDown(delta)
	}
}

func (ui *UI) pageUp(h int) {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveUp(h - 1))
	default:
		ui.view.ScrollUp(h)
	}
}

func (ui *UI) pageDown(h int) {
	switch {
	case ui.state.Mode().IsSelect():
		ui.prompt.SetInput(ui.view.MoveDown(h - 1))
	default:
		ui.view.ScrollDown(h)
	}
}

func (ui *UI) scrollUp(h int, mod tcell.ModMask) {
	m := ui.state.Mode()

	switch {
	case m.IsPrompt() && !m.IsSelect():
		ui.prompt.SetInput(ui.history.PrevLine())
	case mod&tcell.ModShift != 0 && mod&tcell.ModCtrl != 0:
		ui.gotoHome()
	case mod&tcell.ModShift != 0:
		ui.pageUp(h)
	default:
		ui.lineUp()
	}
}

func (ui *UI) scrollDown(h int, mod tcell.ModMask) {
	m := ui.state.Mode()

	switch {
	case m.IsPrompt() && !m.IsSelect():
		ui.prompt.SetInput(ui.history.NextLine())
	case mod&tcell.ModShift != 0 && mod&tcell.ModCtrl != 0:
		ui.gotoEnd()
	case mod&tcell.ModShift != 0:
		ui.pageDown(h)
	default:
		ui.lineDown()
	}
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
	if !ui.state.SwitchMode(m) {
		return
	}

	// former mode
	if ui.state.Last().IsPrompt() {
		ui.prompt.SetInput("")
	}

	// actual mode
	ui.prompt.Lock(!m.IsPrompt())

	// force the cursor off
	if ui.prompt.Locked() {
		ui.state.Root.HideCursor()
	}

	// config completion
	switch m {
	case mode.Open:
		ui.view.LoadPath(ui.state.Path())
		ui.state.SetFindable(fs.Find)
	default:
		ui.state.SetFindable(ui.history.FindLine)
	}

	if ui.state.Last().IsStatic() || m.IsStatic() {
		ui.view.Reset()
	}
}

func (ui *UI) changeBack() {
	if ui.state.Mode() == mode.Chat {
		return // never change chat mode
	}

	if ui.state.Last().IsPrompt() {
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
		log.Println(err)
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

	if flags.Get().Optional.NoPlugins {
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
		return // heap not applicable
	}

	if !ai.Check() {
		ui.overlay.SendError("Assistant is not available")
		return // server not reachable
	}

	title := format(h.String(), "assistant")

	if v, ok := ui.chats.Load(title); !ok {
		c = chat.New(ui.state, h)
		ui.chats.Store(title, c)
	} else {
		c = v.(*chat.Chat)
	}

	path := c.File.Name()

	ui.changeMode(mode.Chat)
	hs.OpenChat(path, path, title)
}
