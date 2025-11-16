package heap

import (
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/cuhsat/fox/v4/internal/pkg/run"
	"github.com/edsrzf/mmap-go"

	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/smap"
)

type Heap struct {
	sync.RWMutex

	Title string // heap title
	Path  string // file path
	Base  string // base path

	Type types.Heap // heap type

	mmap *mmap.MMap // memory map
	smap *smap.SMap // string map

	filters []*Filter // filters

	size int64   // file size
	file fs.File // file handle
}

func New(title, path, base string, ht types.Heap) *Heap {
	heap := &Heap{
		Title: title,
		Path:  path,
		Base:  base,
		Type:  ht,
	}

	return heap
}

func (h *Heap) MMap() *mmap.MMap {
	h.RLock()
	defer h.RUnlock()
	return h.mmap
}

func (h *Heap) SMap() *smap.SMap {
	h.RLock()
	defer h.RUnlock()
	return h.smap
}

func (h *Heap) FMap() *smap.SMap {
	h.RLock()
	defer h.RUnlock()
	return h.LastFilter().fmap
}

func (h *Heap) Size() int64 {
	h.RLock()
	defer h.RUnlock()
	return h.size
}

func (h *Heap) Length() int {
	h.RLock()
	defer h.RUnlock()
	return len(*h.smap)
}

func (h *Heap) String() string {
	switch h.Type {
	case types.Regular:
		return h.Path
	case types.Stdin:
		return "pipe"
	default:
		return h.Title
	}
}

func (h *Heap) Ensure() *Heap {
	if h.file == nil {
		h.Reload()

		// apply global filters once
		if h.Type != types.Chat {
			//			filters := flags.CLI.Filters

			//			for _, filter := range filters.Patterns {
			//				h.AddFilter(
			//					filter,
			//					filters.Before,
			//					filters.After,
			//				)
			//			}
		}
	}

	return h
}

func (h *Heap) Reload(limit run.Limits) {
	var err error

	h.Lock()

	if h.file == nil {
		h.file = fs.Open(h.Path)
	}

	fi, err := h.file.Stat()

	if err != nil {
		log.Println(err)
	}

	h.size = fi.Size()

	if h.mmap != nil {
		_ = h.mmap.Unmap()
	}

	if h.size == 0 {
		h.mmap = new(mmap.MMap) // empty files will cause issues
	} else {
		var m mmap.MMap

		switch f := h.file.(type) {

		// regular file
		case *os.File:
			m, err = mmap.Map(f, mmap.RDONLY, 0)

		// memory file
		case fs.File:
			m, err = fs.Map(f)
		}

		if err != nil {
			log.Println(err)
		}

		h.mmap = &m
	}

	// reduce mmap
	h.mmap = limit.ReduceMMap(h.mmap)

	// reduce smap
	h.smap = limit.ReduceSMap(smap.Map(h.mmap))

	// resets filters
	h.filters = h.filters[:0]
	h.filters = append(h.filters, &Filter{
		new(Pattern),
		new(Context),
		h.smap,
	})

	h.Unlock()

	runtime.GC()
}

func (h *Heap) ThrowAway() {
	h.Lock()

	clear(h.filters)

	h.size = 0
	h.smap = nil

	if h.mmap != nil {
		_ = h.mmap.Unmap()
		h.mmap = nil
	}

	if h.file != nil {
		_ = h.file.Close()
		h.file = nil
	}

	h.Unlock()

	runtime.GC()
}
