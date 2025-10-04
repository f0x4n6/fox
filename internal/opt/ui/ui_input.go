//go:build !minimal

package ui

import (
	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

func (ui *UI) handleMouse(hs *heapset.HeapSet, h *heap.Heap, ev *tcell.EventMouse) {
	if flags.Get().Opt.NoMouse {
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
