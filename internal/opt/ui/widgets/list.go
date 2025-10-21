package widgets

import (
	"sync/atomic"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type List struct {
	base

	h int

	last  int
	delta int

	values atomic.Value
}

func NewList(state *opt.State) *List {
	return &List{
		base: base{state},
	}
}

func (l *List) Render(hs *heapset.HeapSet, x, y, w, h int) int {
	l.h = h - 1 // fill all but the least line

	// debug
	l.SetValues([]string{
		"Hello World 1",
		"Hello World 2",
		"Hello World 3",
		"Hello World 4",
		"Hello World 5",
		"Hello World 6",
		"Hello World 7",
		"Hello World 8",
		"Hello World 9",
		"Hello World 10",
		"Hello World 11",
		"Hello World 12",
		"Hello World 13",
		"Hello World 14",
		"Hello World 15",
		"Hello World 16",
		"Hello World 17",
		"Hello World 18",
		"Hello World 19",
		"Hello World 20",
		"Hello World 21",
		"Hello World 22",
		"Hello World 23",
		"Hello World 24",
		"Hello World 25",
		"Hello World 26",
		"Hello World 27",
		"Hello World 28",
		"Hello World 29",
		"Hello World 30",
	})

	off := max(h-l.delta, 0)

	lines, _ := l.values.Load().([]string)

	for i, line := range lines[off:] {
		l.print(x, y+i, line, themes.Terminal)
	}

	l.print(x, y+l.delta, "Hello World", themes.Subtext2)

	return l.h
}

func (l *List) SetValues(vs []string) {
	l.values.Store(vs)

	l.last = len(vs) - 1
}

func (l *List) GetValue() string {
	return ""
}

func (l *List) MoveUp(delta int) {
	l.delta = max(l.delta-delta, 0)
}

func (l *List) MoveDown(delta int) {
	l.delta = min(l.delta+delta, l.last)
}
