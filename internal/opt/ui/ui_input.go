//go:build !minimal

package ui

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/opt/ai/chat"
	"github.com/cuhsat/fox/internal/opt/ui/widgets"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/mode"
	"github.com/cuhsat/fox/internal/pkg/user/plugins"
)

const (
	delta = 1 // lines
	wheel = 5 // lines
)

func (ui *UI) handleMouse(hs *heapset.HeapSet, h *heap.Heap, ev *tcell.EventMouse) {
	if flags.Get().Optional.NoMouse {
		return
	}

	b := ev.Buttons()

	switch {
	case b&tcell.ButtonMiddle != 0:
		ui.root.GetClipboard()
	case b&tcell.Button4 != 0:
		ui.nextTab(hs, h)
	case b&tcell.Button5 != 0:
		ui.prevTab(hs, h)
	case b&tcell.WheelUp != 0:
		ui.view.ScrollUp(wheel)
	case b&tcell.WheelDown != 0:
		ui.view.ScrollDown(wheel)
	case b&tcell.WheelLeft != 0:
		ui.view.ScrollLeft(wheel)
	case b&tcell.WheelRight != 0:
		ui.view.ScrollRight(wheel)
	}
}

func (ui *UI) handleKey(hs *heapset.HeapSet, h *heap.Heap, ev *tcell.EventKey) bool {
	rootW, rootH := ui.root.Size()

	pageW := rootW - 1 // minus text abbreviation
	pageH := rootH - 2 // minus title and status

	if ui.state.IsNavi() {
		pageW -= text.Dec(h.Length()) + 1
	}

	if ev.Key() != tcell.KeyEscape {
		ui.lkey = ev.Key() // reset
	}

	m := ev.Modifiers()

	switch k := ev.Key(); k {
	case tcell.KeyEscape:
		switch {
		case ui.lkey == tcell.KeyEscape:
			return true // twice to exit
		case ui.state.Mode().IsSelect():
			ui.changeBack()
		case ui.state.Mode().IsPrompt():
			ui.changeBack()
		}

		ui.lkey = k

	case tcell.KeyTab:
		switch {
		case ui.state.Mode().IsSelect():
			return false // stay
		case m&tcell.ModShift != 0:
			ui.prevTab(hs, h)
		default:
			ui.nextTab(hs, h)
		}

	case tcell.KeyF1:
		ui.view.Reset()
		hs.OpenHelp()

	case tcell.KeyF2:
		ui.runAssistant(hs, h)

	case tcell.KeyF3:
		hs.Counts()
		ui.changeMode(mode.Default)

	case tcell.KeyF4:
		hs.Entropy(0.0, 1.0)
		ui.changeMode(mode.Default)

	case tcell.KeyF5:
		hs.Strings(3, math.MaxInt, true, nil)
		ui.changeMode(mode.Default)

	case tcell.KeyF6:
		hs.HashSum(types.SHA256)
		ui.changeMode(mode.Default)

	case tcell.KeyF7:
		hs.Timeline(false)
		ui.changeMode(mode.Default)

	case tcell.KeyF8:
		fallthrough
	case tcell.KeyF9:
		fallthrough
	case tcell.KeyF10:
		fallthrough
	case tcell.KeyF11:
		fallthrough
	case tcell.KeyF12:
		fallthrough
	case tcell.KeyF13:
		fallthrough
	case tcell.KeyF14:
		fallthrough
	case tcell.KeyF15:
		fallthrough
	case tcell.KeyF16:
		fallthrough
	case tcell.KeyF17:
		fallthrough
	case tcell.KeyF18:
		fallthrough
	case tcell.KeyF19:
		fallthrough
	case tcell.KeyF20:
		fallthrough
	case tcell.KeyF21:
		fallthrough
	case tcell.KeyF22:
		fallthrough
	case tcell.KeyF23:
		fallthrough
	case tcell.KeyF24:
		ui.runPlugin(hs, h, strings.ToLower(ev.Name()))

	case tcell.KeyUp:
		ui.scrollUp(pageH, m)

	case tcell.KeyDown:
		ui.scrollDown(pageH, m)

	case tcell.KeyLeft:
		switch {
		case ui.state.Mode().IsPrompt():
			if m&tcell.ModCtrl != 0 {
				ui.prompt.MoveStart()
			} else {
				ui.prompt.Move(-1)
			}
		case m&tcell.ModCtrl != 0:
			ui.prevTab(hs, h)
		case m&tcell.ModShift != 0:
			ui.view.ScrollLeft(pageW)
		default:
			ui.view.ScrollLeft(delta)
		}

	case tcell.KeyRight:
		switch {
		case ui.state.Mode().IsPrompt():
			switch {
			case !ui.prompt.CanMovedEnd():
				ui.prompt.Complete()
			case m&tcell.ModCtrl != 0:
				ui.prompt.MoveEnd()
			default:
				ui.prompt.Move(+1)
			}
		case m&tcell.ModCtrl != 0:
			ui.nextTab(hs, h)
		case m&tcell.ModShift != 0:
			ui.view.ScrollRight(pageW)
		default:
			ui.view.ScrollRight(delta)
		}

	case tcell.KeyHome:
		ui.gotoHome()

	case tcell.KeyPgUp:
		ui.pageUp(pageH)

	case tcell.KeyPgDn:
		ui.pageDown(pageH)

	case tcell.KeyEnd:
		ui.gotoEnd()

	case tcell.KeyCtrlCarat:
		ui.changeTheme()

	case tcell.KeyCtrlSpace:
		fallthrough

	case tcell.KeyCtrlG:
		ui.changeMode(mode.Goto)

	case tcell.KeyCtrlO:
		ui.view.LoadPath(ui.state.Path())
		ui.changeMode(mode.Open)

	case tcell.KeyCtrlL:
		ui.changeMode(mode.Less)

	case tcell.KeyCtrlF:
		ui.changeMode(mode.Grep)

	case tcell.KeyCtrlX:
		ui.changeMode(mode.Hex)

	case tcell.KeyCtrlP:
		ui.changeMode(mode.Pick)

	case tcell.KeyCtrlY:
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
		if h.ModContext(-1) {
			ui.view.Reset()
		}

	case tcell.KeyCtrlK:
		if h.ModContext(+1) {
			ui.view.Reset()
		}

	case tcell.KeyCtrlV:
		ui.root.GetClipboard()

	case tcell.KeyCtrlU:
		if !ui.state.Mode().IsStatic() && hs.Unique() {
			ui.overlay.SendInfo("Union all open files")
		}

	case tcell.KeyCtrlC:
		if !ui.state.Mode().IsStatic() {
			ui.root.SetClipboard(h.Bytes())
			ui.overlay.SendInfo(fmt.Sprintf("%s copied to clipboard", h.String()))
		}

	case tcell.KeyCtrlS:
		if !ui.state.Mode().IsStatic() && ui.bag.Put(h) {
			ui.overlay.SendInfo(fmt.Sprintf("%s saved to %s", h.String(), ui.bag.String()))
		}

	case tcell.KeyCtrlB:
		if fs.Exists(ui.bag.Path) {
			ui.view.Reset()
			hs.OpenFile(ui.bag.Path, ui.bag.Path, ui.bag.Path, types.Ignore)
		} else {
			ui.overlay.SendError(fmt.Sprintf("%s not found", ui.bag.Path))
		}

	case tcell.KeyCtrlQ:
		ui.view.Reset()

		if hs.CloseHeap() == nil {
			return true // exit
		}

	case tcell.KeyCtrlZ:
		ui.runShell()

	case tcell.KeyEnter:
		l := ui.prompt.ReadLine()
		m := ui.state.Mode()

		if m.IsPrompt() && len(strings.TrimSpace(l)) == 0 {
			return false
		}

		ui.history.AddLine(l)

		switch m {
		case mode.Less, mode.Hex:
			ui.view.ScrollLine()

		case mode.Grep:
			ui.view.Reset()
			h.AddFilter(l, 0, 0)
			ui.changeMode(mode.Less)

		case mode.Pick:
			ui.view.Reset()
			h.Select(l, 0, 0)
			ui.changeMode(mode.Less)

		case mode.Goto:
			ui.view.GotoPosition(l)
			ui.changeMode(ui.state.Last())

		case mode.Open:
			if !ui.view.Select() {
				ui.prompt.SetInput(fs.ActualDir)
			} else {
				ui.changeMode(mode.Default)
				ui.state.Call(func() {
					dir := ui.state.Path()

					if l != fs.ActualDir {
						hs.Open(filepath.Join(dir, l))
					} else {
						hs.Open(dir)
					}
				})
			}

		case mode.Chat:
			ui.view.Reset()
			ui.state.Call(func() {
				if v, ok := ui.chats.Load(h.String()); ok {
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
		case len(ui.prompt.GetInput()) > 0:
			ui.prompt.DelRune(widgets.Before)
		case ui.state.Mode().IsSelect():
			ui.changeBack()
		case ui.state.Mode().IsPrompt():
			ui.changeBack()
		case len(h.Patterns()) > 0 && !ui.state.Mode().IsStatic():
			ui.view.Reset()
			h.DelFilter()
		}

	default:
		r := ev.Rune()

		switch r {
		case 0: // error
			return false

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

	return false
}
