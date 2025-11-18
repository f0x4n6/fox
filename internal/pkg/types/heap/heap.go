package heap

import (
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/edsrzf/mmap-go"

	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/smap"
)

type Context struct {
	Name   string
	Type   types.Heap
	Limit  *types.Limits
	Filter *types.Filters
}

type Heap struct {
	sync.RWMutex

	Name string     // heap name
	Type types.Heap // heap type

	mmap mmap.MMap // memory map
	smap smap.SMap // string map

	size int64 // file size
}

func New(ctx *Context, src any) *Heap {
	h := &Heap{
		Name: ctx.Name,
		Type: ctx.Type,
	}

	switch v := src.(type) {
	case []byte:
		h.mmap = src.(mmap.MMap)

	case string:
		f, err := os.Open(v)

		if err != nil {
			log.Println(err)
		}

		fi, err := f.Stat()

		if err != nil {
			log.Println(err)
		}

		h.size = fi.Size()

		// empty files will cause issues
		if h.size == 0 {
			h.mmap = make(mmap.MMap, 0)
		} else {
			h.mmap, err = mmap.Map(f, mmap.RDONLY, 0)

			if err != nil {
				log.Println(err)
			}

			_ = f.Close()
		}

	default:
		log.Fatal("invalid heap type")
	}

	// reduce mmap
	h.mmap = ctx.Limit.ReduceMMap(h.mmap)

	// reduce smap
	h.smap = ctx.Limit.ReduceSMap(smap.Map(h.mmap))

	// filter smap
	h.smap = ctx.Filter.FilterSMap(h.smap)

	return h
}

func (h *Heap) MMap() mmap.MMap {
	h.RLock()
	defer h.RUnlock()
	return h.mmap
}

func (h *Heap) SMap() smap.SMap {
	h.RLock()
	defer h.RUnlock()
	return h.smap
}

func (h *Heap) Size() int64 {
	h.RLock()
	defer h.RUnlock()
	return h.size
}

func (h *Heap) Len() int {
	h.RLock()
	defer h.RUnlock()
	return len(h.smap)
}

func (h *Heap) String() string {
	switch h.Type {
	case types.Stdin:
		return "stdin"
	default:
		return h.Name
	}
}

func (h *Heap) ThrowAway() {
	h.Lock()

	h.size = 0
	//h.smap = nil

	// TODO
	//if h.mmap != nil {
	//	_ = h.mmap.Unmap()
	//	h.mmap = nil
	//}

	h.Unlock()

	runtime.GC()
}
