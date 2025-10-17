package heap

import (
	"log"
	"regexp"
	"strings"

	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/smap"
)

type Pattern struct {
	Value string         // pattern value
	Regex *regexp.Regexp // pattern type
}

type Context struct {
	B    int        // context before
	A    int        // context after
	base *smap.SMap // context base map
}

type Filter struct {
	Pattern *Pattern   // filter pattern
	Context *Context   // filter context
	fmap    *smap.SMap // filter string map
}

func (c *Context) Size() int {
	return (c.B + c.A) / 2
}

func (f *Filter) Len() int {
	if f.Context.base != nil {
		return len(*f.Context.base)
	} else {
		return len(*f.fmap)
	}
}

func (h *Heap) Select(lines string, b, a int) {
	fmap := h.FMap()
	last := h.LastFilter()

	// use only the base of the context for filtering
	if last.Context.base != nil {
		fmap = last.Context.base
	}

	fmap = fmap.Pick(h.parseLines(lines))

	ptn := &Pattern{lines, nil}
	ctx := &Context{b, a, fmap}

	// add global context
	if b+a > 0 {
		fmap = h.addContext(fmap, ctx)
	}

	h.Lock()

	h.filters = append(h.filters, &Filter{
		ptn, ctx, fmap,
	})

	h.Unlock()
}

func (h *Heap) AddFilter(pattern string, b, a int) {
	re, err := regexp.Compile(pattern)

	if err != nil {
		log.Println(err)
		return
	}

	fmap := h.FMap()
	last := h.LastFilter()

	// use only the base of the context for filtering
	if last.Context.base != nil {
		fmap = last.Context.base
	}

	fmap = fmap.Grep(re)

	ptn := &Pattern{pattern, re}
	ctx := &Context{b, a, fmap}

	// add global context
	if b+a > 0 {
		fmap = h.addContext(fmap, ctx)
	}

	h.Lock()

	h.filters = append(h.filters, &Filter{
		ptn, ctx, fmap,
	})

	h.Unlock()
}

func (h *Heap) DelFilter() {
	h.Lock()

	l := len(h.filters)

	if l > 1 {
		h.filters = h.filters[:l-1]
	}

	h.Unlock()
}

func (h *Heap) Filters() []*Filter {
	h.RLock()
	defer h.RUnlock()

	var fs []*Filter

	for _, f := range h.filters[1:] {
		fs = append(fs, f)
	}

	return fs
}

func (h *Heap) Patterns() []*Pattern {
	h.RLock()
	defer h.RUnlock()

	var ps []*Pattern

	for _, f := range h.filters[1:] {
		ps = append(ps, f.Pattern)
	}

	return ps
}

func (h *Heap) LastFilter() *Filter {
	h.RLock()
	defer h.RUnlock()
	return h.filters[max(len(h.filters)-1, 0)]
}

func (h *Heap) HasContext() bool {
	last := h.LastFilter()

	return last.Context.B+last.Context.A > 0
}

func (h *Heap) ModContext(delta int) bool {
	last := h.LastFilter()

	if last.Context.base == nil {
		return false // not filtered
	}

	m := len(*h.SMap())

	// modify current context
	ctx := &Context{
		min(max(last.Context.B+delta, 0), m),
		min(max(last.Context.A+delta, 0), m),
		last.Context.base,
	}

	// readd current context
	fmap := h.addContext(last.Context.base, ctx)

	h.Lock()
	last.Context = ctx
	last.fmap = fmap
	h.Unlock()

	return true
}

func (h *Heap) addContext(s *smap.SMap, ctx *Context) *smap.SMap {
	base := h.SMap()
	fmap := make(smap.SMap, 0, len(*base))

	for grp, str := range *s {
		for _, b := range (*base)[max((str.Nr-1)-ctx.B, 0) : str.Nr-1] {
			b.Grp = grp + 1
			fmap = append(fmap, b)
		}

		str.Grp = grp + 1
		fmap = append(fmap, str)

		for _, a := range (*base)[str.Nr:min(str.Nr+ctx.A, len(*base))] {
			a.Grp = grp + 1
			fmap = append(fmap, a)
		}
	}

	return &fmap
}

func (h *Heap) parseLines(lines string) (nrs []int) {
	for _, l := range strings.Split(lines, ",") {
		n := strings.Split(l, "-")

		if len(n) > 1 {
			a := text.Int(n[0])
			b := text.Int(n[1])

			if a > 0 && b > 0 && a <= b {
				for i := a; i <= b; i++ {
					nrs = append(nrs, i)
				}
			}
		} else {
			if nr := text.Int(l); nr > 0 {
				nrs = append(nrs, nr)
			}
		}
	}

	return
}
