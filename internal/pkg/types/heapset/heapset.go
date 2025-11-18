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

	buf, err := mmap.Map(f, mmap.RDONLY, 0)

	if err != nil {
		log.Println(err)
		return
	}

	///

	if !hs.opts.NoDeflate {
		ok := false

		for {
			buf, ok = hs.deflate(path, buf)

			if !ok {
				break
			}
		}
	}

	if !hs.opts.NoConvert {
		if hs.convert(path, buf) {
			return
		}
	}

	hs.addFile(path, buf)
}

func (hs *HeapSet) deflate(path string, b []byte) ([]byte, bool) {
	var err error

	switch {
	case br.Detect(b):
		b, err = br.Deflate(b)
	case bzip2.Detect(b):
		b, err = bzip2.Deflate(b)
	case gzip.Detect(b):
		b, err = gzip.Deflate(b)
	case lz4.Detect(b):
		b, err = lz4.Deflate(b)
	case xz.Detect(b):
		b, err = xz.Deflate(b)
	case zlib.Detect(b):
		b, err = zlib.Deflate(b)
	case zstd.Detect(b):
		b, err = zstd.Deflate(b)
	default:

	}

	if err != nil {
		log.Println(err)
		return true
	}

	if len(b) > 0 {
		hs.addData(path, b)
		return true
	}

	switch {
	case rar.Detect(b):
		hs.extract(path, rar.Extract)
		return true
	case tar.Detect(b):
		hs.extract(path, tar.Extract)
		return true
	case zip.Detect(b):
		hs.extract(path, zip.Extract)
		return true
	}

	return false
}

func (hs *HeapSet) extract(path string, fn files.Extract) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("corrupted archive or wrong password")
			return
		}
	}()

	items := fn(path, hs.opts.Password)

	if len(items) == 0 {
		panic("no item(s)")
	}

	for _, i := range items {
		i.Path = hs.deflate(i.Path)

		if len(i.Path) == 0 {
			continue
		}

		i.Path = hs.convert(i.Path)

		if len(i.Path) == 0 {
			continue
		}

		hs.addData(i.Path)
	}
}

func (hs *HeapSet) convert(path string, b []byte) bool {
	switch {
	case evtx.Detect(b):
		if buf, err := evtx.Convert(b); err == nil {
			hs.addData(path, buf)
		} else {
			log.Println(err)
		}

		return true

	case journal.Detect(b):
		if buf, err := journal.Convert(b); err == nil {
			hs.addData(path, buf)
		} else {
			log.Println(err)
		}

		return true
	}

	return false
}

func (hs *HeapSet) addPipe() {
	buf, err := sys.Stdin()

	if err != nil {
		log.Fatal(err)
	}

	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   Stdin,
			Type:   types.Stdin,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, buf,
	))
}

func (hs *HeapSet) addFile(path string, b []byte) {
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   path,
			Type:   types.Regular,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, b,
	))
}

func (hs *HeapSet) addData(name string, b []byte) {
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   name,
			Type:   types.Deflate,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, b,
	))
}
