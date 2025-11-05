package widgets

import (
	"sync/atomic"

	"github.com/cuhsat/fox/v3/internal/opt"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v3/internal/pkg/types/mode"
	"github.com/cuhsat/fox/v3/internal/pkg/types/smap"
)

type View struct {
	base
	heap  *heap.Heap
	fmap  *smap.SMap
	list  atomic.Value
	cache map[string]position
	trans map[int]int

	nr int

	off  int
	line int

	last  point
	delta point
	plane plane
}

type position struct {
	nr    int
	len   int
	delta point
}

func NewView(state *opt.State) *View {
	return &View{
		base:  base{state},
		cache: make(map[string]position),
	}
}

func (v *View) Render(hs *heapset.HeapSet, x, y, w, h int) int {
	v.plane = plane{x, y, w, h - 1} // fill all but the least line

	if hs != nil {
		_, v.heap = hs.Heap()
	} else {
		return v.plane.H
	}

	switch v.state.Mode() {
	case mode.Hex:
		v.hexRender(&v.plane)
	case mode.Open:
		v.listRender(&v.plane)
	default:
		v.textRender(&v.plane)
	}

	return v.plane.H
}

func (v *View) Reset() {
	v.delta.X = 0
	v.delta.Y = 0

	v.off = 0

	v.nr = 0
}

func (v *View) Select() bool {
	item := v.listItem()

	if item.Text == fs.ActualDir {
		return true // select root
	}

	if item.Leaf {
		return true // select line
	}

	v.LoadPath(item.Value)

	return false
}

func (v *View) Preserve() {
	if v.fmap != nil && len(*v.fmap) > v.delta.Y {
		v.nr = (*v.fmap)[v.delta.Y].Nr
	}
}

func (v *View) GotoPosition(pos string) {
	if !v.state.Mode().IsStatic() {
		v.textGoto(pos)
	}
}

func (v *View) SavePosition(key string) {
	if v.fmap != nil && len(*v.fmap) > v.delta.Y {
		v.cache[key] = position{
			nr:  (*v.fmap)[v.delta.Y].Nr,
			len: len(*v.fmap),
			delta: point{
				v.delta.X,
				v.delta.Y,
			},
		}
	}
}

func (v *View) LoadPosition(key string) {
	// safe defaults
	v.delta.X = 0
	v.delta.Y = 0

	// can be nil in hex mode
	if s, ok := v.cache[key]; ok && v.fmap != nil {
		if s.len == len(*v.fmap) {
			v.delta = s.delta
		} else {
			v.nr = s.nr
		}
	}
}

func (v *View) LoadPath(root string) {
	items := v.state.List(root)

	v.state.ChangePath(root)

	v.list.Store(items)
	v.last = point{0, len(items) - 1}
	v.line = 0
	v.off = 0
}

func (v *View) MarkLine(x, y int, z bool) {
	// check horizontal out-of-bounds
	if x < v.plane.X || x > v.plane.X+v.plane.W {
		return
	}

	// check vertical out-of-bounds
	if y < v.plane.Y || y > v.plane.Y+v.plane.H {
		return
	}

	if !v.state.Mode().IsStatic() {
		v.textMark(x-v.plane.X, y-v.plane.Y, z)
	}
}

func (v *View) MoveHome() string {
	v.line, v.off = 0, 0

	return v.listItem().Text
}

func (v *View) MoveEnd() string {
	v.line = min(v.last.Y, v.plane.H-1)
	v.off = max((v.last.Y-v.plane.H)+1, 0)

	return v.listItem().Text
}

func (v *View) MoveUp(delta int) string {
	if v.line > 0 {
		v.line = max(v.line-delta, 0)
	} else if v.off > 0 {
		v.off = max(v.off-delta, 0)
	}

	return v.listItem().Text
}

func (v *View) MoveDown(delta int) string {
	if v.line < v.plane.H-1 {
		v.line = min(v.line+delta, v.plane.H-1, v.last.Y)
	} else if v.off < (v.last.Y-v.plane.H)+1 {
		v.off = min(v.off+delta, (v.last.Y-v.plane.H)+1)
	}

	return v.listItem().Text
}

func (v *View) ScrollLine() {
	if v.state.Mode().IsStatic() || v.heap.HasContext() {
		v.ScrollDown(1)
		return
	}

	if v.fmap == nil || len(*v.fmap) <= 1 {
		return
	}

	v.nr = (*v.fmap)[v.delta.Y].Nr

	for y := v.delta.Y; y < len(*v.fmap); y++ {
		if v.nr < (*v.fmap)[y].Nr {
			v.nr = (*v.fmap)[y].Nr
			break
		}
	}
}

func (v *View) ScrollLast() {
	v.delta.Y = max(v.last.Y-(v.plane.H-3), 0)
}

func (v *View) ScrollHome() {
	v.delta.Y = 0
}

func (v *View) ScrollEnd() {
	v.delta.Y = v.last.Y
}

func (v *View) ScrollTo(x, y int) {
	v.delta.X = max(min(x, v.last.X), 0)
	v.delta.Y = max(min(y, v.last.Y), 0)
}

func (v *View) ScrollUp(delta int) {
	v.delta.Y = max(v.delta.Y-delta, 0)
}

func (v *View) ScrollDown(delta int) {
	v.delta.Y = min(v.delta.Y+delta, v.last.Y)
}

func (v *View) ScrollLeft(delta int) {
	v.delta.X = max(v.delta.X-delta, 0)
}

func (v *View) ScrollRight(delta int) {
	v.delta.X = min(v.delta.X+delta, v.last.X)
}
