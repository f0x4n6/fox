package heapset

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/rar"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/tar"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/zip"
	"github.com/cuhsat/fox/v4/internal/pkg/files/convert/evtx"
	"github.com/cuhsat/fox/v4/internal/pkg/files/convert/journal"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/br"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/bzip2"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/gzip"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/lz4"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/xz"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/zlib"
	"github.com/cuhsat/fox/v4/internal/pkg/files/deflate/zstd"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/edsrzf/mmap-go"
)

const Stdin = "-"

type Options struct {
	Limit     *types.Limits
	Filter    *types.Filters
	Password  string
	NoDeflate bool
	NoConvert bool
}

type HeapSet struct {
	sync.RWMutex
	opts  *Options     // set options
	heaps []*heap.Heap // set heaps
}

func New(paths []string, opts *Options) *HeapSet {
	hs := HeapSet{opts: opts}

	if sys.Piped(os.Stdin) {
		paths = append(paths, Stdin)
	}

	for _, path := range paths {
		if path == Stdin {
			hs.addPipe()
			return &hs
		}

		_, err := os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			log.Printf("%s does not exist\n", path)
		}

		hs.loadPath(path)
	}

	//  TODO?
	//	if len(hs.heaps) == 0 {
	//		log.Fatal("could not load any files")
	//	}

	return &hs
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

func (hs *HeapSet) loadPath(path string) {
	match, err := doublestar.FilepathGlob(path)

	if err != nil {
		log.Println(err)
		return
	}

	for _, m := range match {
		fi, err := os.Stat(m)

		if err != nil {
			log.Println(err)
			continue
		}

		if fi.IsDir() {
			hs.loadDir(m)
		} else {
			hs.loadFile(m)
		}
	}
}

func (hs *HeapSet) loadDir(path string) {
	dir, err := os.ReadDir(path)

	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range dir {
		if !f.IsDir() {
			hs.loadFile(filepath.Join(path, f.Name()))
		}
	}
}

func (hs *HeapSet) loadFile(path string) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0x400)

	if err != nil {
		log.Println(err)
		return
	}

	defer sys.Handle(f.Close)

	fi, err := f.Stat()

	if err != nil {
		log.Println(err)
		return
	}

	// empty files will cause issues
	if fi.Size() == 0 {
		hs.addData(path, []byte{})
		return
	}

	b, err := mmap.Map(f, mmap.RDONLY, 0)

	if err != nil {
		log.Println(err)
		return
	}

	hs.process(path, b, false)
}

func (hs *HeapSet) process(path string, b []byte, data bool) {
	var ok bool

	if !hs.opts.NoDeflate {
		if b, ok = hs.deflate(b); ok {
			data = true
		}

		if hs.extract(path, b) {
			return
		}
	}

	if !hs.opts.NoConvert {
		if b, ok = hs.convert(b); ok {
			data = true
		}
	}

	if data {
		hs.addData(path, b)
	} else {
		hs.addFile(path, b)
	}
}

func (hs *HeapSet) extract(path string, b []byte) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("archive corrupt or password wrong")
			return
		}
	}()

	var fn files.Extract

	switch {
	case rar.Detect(b):
		fn = rar.Extract
	case tar.Detect(b):
		fn = tar.Extract
	case zip.Detect(b):
		fn = zip.Extract
	default:
		return false
	}

	for _, e := range fn(b, path, hs.opts.Password) {
		hs.process(e.Path, e.Data, true)
	}

	return true
}

func (hs *HeapSet) deflate(b []byte) ([]byte, bool) {
	switch {
	case br.Detect(b):
		return call(b, br.Deflate)
	case bzip2.Detect(b):
		return call(b, bzip2.Deflate)
	case gzip.Detect(b):
		return call(b, gzip.Deflate)
	case lz4.Detect(b):
		return call(b, lz4.Deflate)
	case xz.Detect(b):
		return call(b, xz.Deflate)
	case zlib.Detect(b):
		return call(b, zlib.Deflate)
	case zstd.Detect(b):
		return call(b, zstd.Deflate)
	}

	return b, false
}

func (hs *HeapSet) convert(b []byte) ([]byte, bool) {
	switch {
	case evtx.Detect(b):
		return call(b, evtx.Convert)
	case journal.Detect(b):
		return call(b, journal.Convert)
	}

	return b, false
}

func (hs *HeapSet) addPipe() {
	buf, err := sys.Stdin()

	if err != nil {
		log.Fatal(err)
	}

	hs.Lock()
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   Stdin,
			Type:   types.Stdin,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, buf,
	))
	hs.Unlock()
}

func (hs *HeapSet) addFile(path string, b []byte) {
	hs.Lock()
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   path,
			Type:   types.Regular,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, b,
	))
	hs.Unlock()
}

func (hs *HeapSet) addData(name string, b []byte) {
	hs.Lock()
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   name,
			Type:   types.Deflate,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, b,
	))
	hs.Unlock()
}

func call(b []byte, fn func([]byte) ([]byte, error)) ([]byte, bool) {
	if r, err := fn(b); err == nil {
		return r, true
	} else {
		log.Println(err)
		return b, false
	}
}
