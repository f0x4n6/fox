package heapset

import (
	"cmp"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
	"sync/atomic"

	"github.com/cuhsat/fox/internal/pkg/files/events"
	"github.com/cuhsat/fox/internal/pkg/files/events/evtx"
	"github.com/cuhsat/fox/internal/pkg/files/events/journal"
	"github.com/cuhsat/fox/internal/pkg/files/timeline/cef"
	"github.com/cuhsat/fox/internal/pkg/files/timeline/plain"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
)

type util func(h *heap.Heap) string

func (hs *HeapSet) Merge() bool {
	var heaps []*heap.Heap

	f := fs.Create("/fox/merged")

	hs.Range(func(i int, h *heap.Heap) bool {
		switch h.Type {
		case types.Regular, types.Deflate:
			_, _ = f.Write(h.Bytes())
			_, _ = f.WriteString("\n")
			h.ThrowAway()

		default:
			heaps = append(heaps, h)
		}
		return true
	})

	fi, _ := f.Stat()

	if fi.Size() == 0 {
		return false
	}

	hs.OpenFile(f.Name(), f.Name(), "merged", types.Ignore)

	return true
}

func (hs *HeapSet) Compare(git bool) *HeapSet {
	var heaps [2]*heap.Heap

	hs.Range(func(i int, h *heap.Heap) bool {
		if h.Type == types.Regular || h.Type == types.Deflate {
			heaps[i] = h
		}
		return true
	})

	f := fs.Create("/fox/compare")

	_, _ = f.WriteString(text.Diff(
		heaps[0].String(),
		heaps[1].String(),
		heaps[0].SMap().Lines(),
		heaps[1].SMap().Lines(),
		git,
	))

	hs.OpenFile(f.Name(), f.Name(), "Compare", types.Stdout)

	return hs
}

func (hs *HeapSet) Counts() *HeapSet {
	hs.reduce("counts", func(h *heap.Heap) string {
		return fmt.Sprintf("%8dL %8dB  %s\n", h.Count(), h.Len(), h.String())
	})

	return hs
}

func (hs *HeapSet) Entropy(n, m float64) *HeapSet {
	hs.reduce("entropy", func(h *heap.Heap) string {
		v := h.Entropy(n, m)

		if v == -1 {
			return "" // filtered
		}

		return fmt.Sprintf("%.10f  %s\n", v, h.String())
	})

	return hs
}

func (hs *HeapSet) HashSum(algos ...string) *HeapSet {
	for _, algo := range algos {
		hs.reduce(algo, func(h *heap.Heap) string {
			sum, err := h.HashSum(algo)

			if err != nil {
				log.Println(err)
			}

			switch algo {
			case types.SDHASH:
				return fmt.Sprintf("%s  %s\n", sum, h.String())
			default:
				return fmt.Sprintf("%x  %s\n", sum, h.String())
			}
		})
	}

	return hs
}

func (hs *HeapSet) Strings(n, m int, i bool, re *regexp.Regexp) *HeapSet {
	hs.reduce("strings", func(h *heap.Heap) string {
		var sb strings.Builder

		for v := range h.Strings(n, m, i, re) {
			sb.WriteString(strings.TrimSpace(v.Str))
			sb.WriteRune('\n')
		}

		return sb.String()
	})

	return hs
}

func (hs *HeapSet) Timeline(c bool) *HeapSet {
	f := fs.Create("/fox/timeline")

	for _, l := range hs.Extract(c) {
		_, _ = f.WriteString(l)
		_, _ = f.WriteString("\n")
	}

	hs.OpenFile(f.Name(), f.Name(), "timeline", types.Stdout)

	return hs
}

func (hs *HeapSet) Extract(c bool) (ls []string) {
	var fn = plain.Format

	if c {
		fn = cef.Format
	}

	hs.Range(func(_ int, h *heap.Heap) bool {
		for _, str := range *h.FMap() {
			for _, ex := range []events.Extract{
				evtx.Extract,
				journal.Extract,
			} {
				if e := ex(str.Str); e != nil {
					ls = append(ls, fn(e))
				}
			}
		}
		return true
	})

	slices.SortStableFunc(ls, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return ls
}

func (hs *HeapSet) reduce(t string, fn util) {
	f := fs.Create(fmt.Sprintf("/fox/%s", t))

	hs.Range(func(i int, h *heap.Heap) bool {
		if h.Type == types.Regular || h.Type == types.Deflate {
			_, err := f.WriteString(fn(h))

			if err != nil {
				log.Println(err)
				return false
			}
		}
		return true
	})

	_ = f.Close()

	if idx, ok := hs.findByName(t); !ok {
		hs.OpenFile(f.Name(), f.Name(), t, types.Stdout)
	} else {
		h := hs.atomicGet(idx)
		h.Path = f.Name()
		h.Reload()

		atomic.StoreInt32(hs.index, idx)
	}
}
