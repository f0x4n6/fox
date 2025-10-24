package widgets

import (
	"fmt"
	"sync/atomic"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ui/adapter"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type List struct {
	base

	adapter adapter.Adapter
	fn      adapter.Callback

	h int

	off  int
	last int
	line int

	values atomic.Value
}

func NewList(state *opt.State) *List {
	return &List{
		base: base{state},
	}
}

func (l *List) Render(hs *heapset.HeapSet, x, y, w, h int) int {
	l.h = h - 1 // fill all but the least line

	page := l.page()

	for i, node := range page {
		if !node.Leaf {
			l.print(x, y+i, fmt.Sprintf("> %s", node.Text), themes.Terminal)
		} else {
			l.print(x, y+i, fmt.Sprintf("  %s", node.Text), themes.Terminal)
		}
	}

	l.print(x+2, y+(l.line), page[l.line].Text, themes.Subtext2)

	return l.h
}

func (l *List) SetAdapter(adapter adapter.Adapter) {
	l.adapter = adapter
}

func (l *List) SetSelected(fn adapter.Callback) {
	l.fn = fn
}

func (l *List) GetValue() string {
	vs := l.values.Load().([]string)

	return vs[l.line]
}

func (l *List) Reset() {
	nodes := l.adapter.Init()

	l.values.Store(nodes)

	l.last = len(nodes) - 1
}

func (l *List) Select() bool {
	page := l.page()
	line := page[l.line]

	v := line.Value

	if !line.Leaf {
		nodes := l.adapter.List(v)

		l.values.Store(nodes)
		l.last = len(nodes) - 1
		l.line = 0
		l.off = 0

		return false
	} else {
		l.state.Call(func() { l.fn(v) })

		return true
	}
}

func (l *List) MoveUp(delta int) {
	if l.line > 0 {
		l.line = max(l.line-delta, 0)
	} else if l.off > 0 {
		l.off = max(l.off-delta, 0)
	}
}

func (l *List) MoveDown(delta int) {
	if l.line < l.h-1 {
		l.line = min(l.line+delta, l.last)
	} else if l.off <= (l.last - l.h) {
		l.off = min(l.off+delta, l.last)
	}
}

func (l *List) page() []adapter.Node {
	nodes := l.values.Load().([]adapter.Node)

	return nodes[l.off:]
}
