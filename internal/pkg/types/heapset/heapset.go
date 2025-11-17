package heapset

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/rar"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/tar"
	"github.com/cuhsat/fox/v4/internal/pkg/files/archive/zip"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/br"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/bzip2"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/gzip"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/lz4"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/xz"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/zlib"
	"github.com/cuhsat/fox/v4/internal/pkg/files/compress/zstd"
	"github.com/cuhsat/fox/v4/internal/pkg/files/parser/evtx"
	"github.com/cuhsat/fox/v4/internal/pkg/files/parser/journal"
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
	hs.load(paths)

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

func (hs *HeapSet) load(paths []string) {
	if sys.Piped(os.Stdin) {
		paths = append(paths, Stdin)
	}

	for _, path := range paths {
		if path == Stdin {
			hs.addPipe()
			break
		}

		_, err := os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			log.Println(fmt.Errorf("%s does not exist", path))
		}

		hs.loadPath(path)
	}

	if len(hs.heaps) == 0 && len(paths) > 0 {
		sys.Panic("could not load any files")
	}

	for _, h := range hs.heaps {
		h.Load(
			hs.opts.Limit,
			hs.opts.Filter,
		)
	}
}

func (hs *HeapSet) loadPath(path string) {
	match, err := doublestar.FilepathGlob(path)

	if err != nil {
		log.Println(err)
	}

	for _, m := range match {
		fi, err := os.Stat(m)

		if err != nil {
			log.Println(err)
			return
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
	base := path

	if !hs.opts.NoDeflate {
		path = hs.deflate(path, base)

		if len(path) == 0 {
			return
		}
	}

	path = hs.process(path)

	if len(path) == 0 {
		return
	}

	hs.addFile(path, base)
}

func (hs *HeapSet) loadArchive(path, base string, fn files.Deflate) {
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
		i.Path = hs.deflate(i.Path, base)

		if len(i.Path) == 0 {
			continue
		}

		i.Path = hs.process(i.Path)

		if len(i.Path) == 0 {
			continue
		}

		hs.addItem(i.Path, base)
	}
}

func (hs *HeapSet) addPipe() {
	pipe := sys.Stdin().Name()

	hs.heaps = append(hs.heaps, heap.New(
		"-",
		pipe,
		pipe,
		types.Stdin,
	))
}

func (hs *HeapSet) addFile(path, base string) {
	var t = types.Regular

	if path != base {
		t = types.Deflate
	}

	hs.heaps = append(hs.heaps, heap.New(
		base,
		path,
		base,
		t,
	))
}

func (hs *HeapSet) addItem(path, base string) {
	hs.heaps = append(hs.heaps, heap.New(
		path,
		path,
		base,
		types.Deflate,
	))
}

func (hs *HeapSet) deflate(path, base string) string {
	// check for compression
	switch {
	case br.Detect(path):
		path = br.Deflate(path)
	case bzip2.Detect(path):
		path = bzip2.Deflate(path)
	case gzip.Detect(path):
		path = gzip.Deflate(path)
	case lz4.Detect(path):
		path = lz4.Deflate(path)
	case xz.Detect(path):
		path = xz.Deflate(path)
	case zlib.Detect(path):
		path = zlib.Deflate(path)
	case zstd.Detect(path):
		path = zstd.Deflate(path)
	}

	// check for archive
	switch {
	case rar.Detect(path):
		hs.loadArchive(path, base, rar.Deflate)
		return ""
	case tar.Detect(path):
		hs.loadArchive(path, base, tar.Deflate)
		return ""
	case zip.Detect(path):
		hs.loadArchive(path, base, zip.Deflate)
		return ""
	}

	return path
}

func (hs *HeapSet) process(path string) string {
	if !hs.opts.NoConvert {
		if evtx.Detect(path) {
			path = evtx.Parse(path)
		}

		if journal.Detect(path) {
			path = journal.Parse(path)
		}
	}

	return path
}
