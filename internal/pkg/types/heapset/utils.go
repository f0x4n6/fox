package heapset

import (
	"fmt"
	"log"
	"strings"
	"sync/atomic"

	"github.com/zeebo/xxh3"

	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type util func(h *heap.Heap) string

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

func (hs *HeapSet) Unique() *HeapSet {
	var lines = make(map[uint64]types.Null)

	hs.reduce("unique", func(h *heap.Heap) string {
		var sb strings.Builder

		for _, s := range *h.SMap() {
			x := xxh3.HashString(s.Str)

			if _, ok := lines[x]; !ok {
				sb.WriteString(s.Str)
				sb.WriteRune('\n')
				lines[x] = types.Null{}
			}
		}

		return sb.String()
	})

	return hs
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
