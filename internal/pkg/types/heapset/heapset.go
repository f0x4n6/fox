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
	if !hs.opts.NoDeflate {
		hs.deflate(path)
	}

	if !hs.convert(path) {
		hs.addFile(path)
	}
}

func (hs *HeapSet) loadArchive(path string, fn files.Deflate) {
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

func (hs *HeapSet) deflate(path string) bool {
	var buf []byte
	var err error

	switch {
	case br.Detect(path):
		buf, err = br.Deflate(path)
	case bzip2.Detect(path):
		buf, err = bzip2.Deflate(path)
	case gzip.Detect(path):
		buf, err = gzip.Deflate(path)
	case lz4.Detect(path):
		buf, err = lz4.Deflate(path)
	case xz.Detect(path):
		buf, err = xz.Deflate(path)
	case zlib.Detect(path):
		buf, err = zlib.Deflate(path)
	case zstd.Detect(path):
		buf, err = zstd.Deflate(path)
	}

	if err != nil {
		log.Println(err)
		return true
	}

	if len(buf) > 0 {
		hs.addData(path, buf)
		return true
	}

	switch {
	case rar.Detect(path):
		hs.loadArchive(path, rar.Deflate)
		return true
	case tar.Detect(path):
		hs.loadArchive(path, tar.Deflate)
		return true
	case zip.Detect(path):
		hs.loadArchive(path, zip.Deflate)
		return true
	}

	return false
}

func (hs *HeapSet) convert(path string) bool {
	if !hs.opts.NoConvert {
		switch {
		case evtx.Detect(path):
			if buf, err := evtx.Convert(path); err == nil {
				hs.addData(path, buf)
			} else {
				log.Println(err)
			}

			return true

		case journal.Detect(path):
			if buf, err := journal.Convert(path); err == nil {
				hs.addData(path, buf)
			} else {
				log.Println(err)
			}

			return true
		}
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

func (hs *HeapSet) addFile(path string) {
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   path,
			Type:   types.Regular,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, path,
	))
}

func (hs *HeapSet) addData(name string, buf []byte) {
	hs.heaps = append(hs.heaps, heap.New(
		&heap.Context{
			Name:   name,
			Type:   types.Deflate,
			Limit:  hs.opts.Limit,
			Filter: hs.opts.Filter,
		}, buf,
	))
}
