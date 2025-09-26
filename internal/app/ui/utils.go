package ui

import (
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

	if h.Type == types.Agent {
		ui.change(mode.Chat)
	} else if ui.ctx.Mode() == mode.Chat {
		ui.change(ui.ctx.Last())
	}
}
