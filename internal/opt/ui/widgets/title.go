package widgets

import (
	"fmt"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type Title struct {
	base
}

func NewTitle(state *opt.State) *Title {
	return &Title{base{state}}
}

func (t *Title) Render(hs *heapset.HeapSet, x, y, w, _ int) int {
	var i int32
	var n int32
	var h *heap.Heap
	var s = "Loading…"

	if hs != nil {
		i, h = hs.Heap()
		s = h.String()
		n = hs.Len()
	}

	var c string

	if n > 1 {
		c = fmt.Sprintf(" %d %c %d ", i, t.state.Icon.VSep, n)
	}

	// render blank line
	t.blank(x, y, w, themes.Surface0)

	// render heap filepath
	t.print(x, y, text.Abr(s, w-(x+text.Len(c)+1)), themes.Surface2)

	// render heapset index and count
	t.print(x+w-text.Len(c), y, c, themes.Surface1)

	return 1
}
