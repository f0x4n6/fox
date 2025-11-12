package heapset

import (
	"slices"
	"sync"
	"sync/atomic"

	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/types/loader"
)

type Changed func(*heap.Heap)

type Range func(int, *heap.Heap) bool

type HeapSet struct {
	sync.RWMutex

	loader *loader.Loader // file loader

	changed Changed // file changed

	heaps []*heap.Heap // set heaps
	index *int32       // set index
}

func New(paths []string) *HeapSet {
	hs := HeapSet{
		loader: loader.New(),
		index:  new(int32),
	}

	go hs.notify()

	for _, h := range hs.loader.Load(paths) {
		hs.atomicAdd(h)
	}

	// load first heap
	hs.LoadHeap()

	return &hs
}

func (hs *HeapSet) Len() int32 {
	hs.RLock()
	defer hs.RUnlock()
	return int32(len(hs.heaps))
}

func (hs *HeapSet) Range(fn Range) {
	hs.RLock()

	for i, h := range hs.heaps {
		if !fn(i, h.Ensure()) {
			break
		}
	}

	hs.RUnlock()
}

func (hs *HeapSet) OpenFile(path, base, title string, tp types.Heap) {
	if !fs.Exists(path) {
		return
	}

	idx, ok := hs.findByPath(path)

	if !ok {
		idx = hs.Len()

		hs.atomicAdd(heap.New(title, path, base, tp))
	}

	atomic.StoreInt32(hs.index, idx)

	hs.LoadHeap()
}

func (hs *HeapSet) LoadHeap() *heap.Heap {
	h := hs.atomicGet(atomic.LoadInt32(hs.index))

	if h.Ensure().Type == types.Regular {
		hs.watchFile(h.Path) // changed file
	}

	return h
}

func (hs *HeapSet) CloseOther() {
	hs.RLock()

	var v []*heap.Heap

	for _, h := range hs.heaps {
		if h.Type == types.Stdout {
			v = append(v, h)
		} else {
			h.ThrowAway()
		}
	}

	hs.heaps = v

	hs.RUnlock()

	atomic.StoreInt32(hs.index, 0)
}

func (hs *HeapSet) ThrowAway() {
	hs.Lock()

	for _, h := range hs.heaps {
		h.ThrowAway()
	}

	hs.heaps = hs.heaps[:0]

	hs.Unlock()

	atomic.AddInt32(hs.index, -1)
}

func (hs *HeapSet) findByPath(path string) (int32, bool) {
	hs.RLock()
	defer hs.RUnlock()

	for i, h := range hs.heaps {
		if h.Base == path {
			return int32(i), true
		}
	}

	return 0, false
}

func (hs *HeapSet) findByName(name string) (int32, bool) {
	hs.RLock()
	defer hs.RUnlock()

	for i, h := range hs.heaps {
		if h.Title == name {
			return int32(i), true
		}
	}

	return 0, false
}

func (hs *HeapSet) atomicAdd(h *heap.Heap) {
	hs.Lock()
	hs.heaps = append(hs.heaps, h)
	hs.Unlock()
}

func (hs *HeapSet) atomicGet(idx int32) *heap.Heap {
	hs.RLock()
	defer hs.RUnlock()
	return hs.heaps[idx]
}

func (hs *HeapSet) atomicDel(idx int32) {
	hs.Lock()
	hs.heaps = slices.Delete(hs.heaps, int(idx), int(idx)+1)
	hs.Unlock()
}
