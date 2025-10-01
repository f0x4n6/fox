//go:build !minimal

package ui

import (
	"github.com/cuhsat/fox/internal/app/ai"
	"github.com/cuhsat/fox/internal/pkg/flags"
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

	ui.view.SaveState(h.Path)

	switch d {
	case Prev:
		h = hs.PrevHeap()
	case Next:
		h = hs.NextHeap()
	}

	ui.view.LoadState(h.Path)

	if h.Type == types.Chat {
		ui.changeMode(mode.Chat)
	} else if ui.ctx.Mode() == mode.Chat {
		ui.changeMode(ui.ctx.Last())
	}
}

func (ui *UI) changeMode(m mode.Mode) {
	// check for examiner support
	if m == mode.Chat && !ai.Check() {
		ui.overlay.SendError("Assistant is not available")
		return
	}

	if !ui.ctx.SwitchMode(m) {
		return
	}

	// former mode
	if ui.ctx.Last().Prompt() {
		ui.prompt.SetValue("")
	}

	// actual mode
	ui.prompt.Lock(!m.Prompt())

	// force the cursor off
	if ui.prompt.Locked() {
		ui.ctx.Root.HideCursor()
	}

	if ui.ctx.Last().Static() || m.Static() {
		ui.view.Reset()
	}
}
