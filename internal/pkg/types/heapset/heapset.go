package heapset

import (
	"sync"

	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/types/loader"
)

type HeapSet struct {
	sync.RWMutex
	heaps []*heap.Heap // set heaps
}

func New(paths []string, opts *loader.Options) *HeapSet {
	heaps := loader.New(opts).Load(paths)

	for _, h := range heaps {
		h.Load(
			opts.Limit,
			opts.Filter,
		)
	}

	return &HeapSet{
		heaps: heaps,
	}
}

func (hs *HeapSet) Len() int {
	hs.RLock()
	defer hs.RUnlock()
	return len(hs.heaps)
}

func (hs *HeapSet) Get() []*heap.Heap {
	hs.RLock()
	defer hs.RUnlock()
	return hs.heaps[:]
}

func (hs *HeapSet) ThrowAway() {
	hs.Lock()

	for _, h := range hs.heaps {
		h.ThrowAway()
	}

	hs.heaps = hs.heaps[:0]

	hs.Unlock()
}
