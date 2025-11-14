package loader

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	"github.com/cuhsat/fox/v4/internal/pkg/flags"
	"github.com/cuhsat/fox/v4/internal/pkg/sys"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

const Stdin = "-"

type Loader struct {
	heaps []*heap.Heap // temp heaps
}

func New() *Loader {
	return new(Loader)
}

func (l *Loader) Load(paths []string) []*heap.Heap {
	if sys.Piped(os.Stdin) {
		paths = append(paths, Stdin)
	}

	for _, path := range paths {
		if path == Stdin {
			l.addPipe()
			break
		}

		_, err := os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			log.Println(fmt.Errorf("%s does not exist", path))
		}

		l.loadPath(path)
	}

	if len(l.heaps) == 0 && len(paths) > 0 {
		sys.Exit("could not load any files")
	}

	return l.heaps
}

func (l *Loader) loadPath(path string) {
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
			l.loadDir(m)
		} else {
			l.loadFile(m)
		}
	}
}

func (l *Loader) loadDir(path string) {
	dir, err := os.ReadDir(path)

	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range dir {
		if !f.IsDir() {
			l.loadFile(filepath.Join(path, f.Name()))
		}
	}
}

func (l *Loader) loadFile(path string) {
	base := path

	if !flags.CLI.NoDeflate {
		path = l.deflate(path, base)

		if len(path) == 0 {
			return
		}
	}

	path = l.process(path)

	if len(path) == 0 {
		return
	}

	l.addFile(path, base)
}

func (l *Loader) loadArchive(path, base string, fn files.Deflate) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("corrupted archive or wrong password")
			return
		}
	}()

	items := fn(path, flags.CLI.Pass)

	if len(items) == 0 {
		panic("no item(s)")
	}

	for _, i := range items {
		i.Path = l.deflate(i.Path, base)

		if len(i.Path) == 0 {
			continue
		}

		i.Path = l.process(i.Path)

		if len(i.Path) == 0 {
			continue
		}

		l.addItem(i.Path, base)
	}
}

func (l *Loader) addPipe() {
	pipe := sys.Stdin().Name()

	l.heaps = append(l.heaps, heap.New(
		"-",
		pipe,
		pipe,
		types.Stdin,
	))
}

func (l *Loader) addFile(path, base string) {
	var t = types.Regular

	if path != base {
		t = types.Deflate
	}

	l.heaps = append(l.heaps, heap.New(
		base,
		path,
		base,
		t,
	))
}

func (l *Loader) addItem(path, base string) {
	l.heaps = append(l.heaps, heap.New(
		path,
		path,
		base,
		types.Deflate,
	))
}

func (l *Loader) addPlugin(path, name string) {
	l.heaps = append(l.heaps, heap.New(
		name,
		path,
		path,
		types.Plugin,
	))
}

func (l *Loader) deflate(path, base string) string {
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
		l.loadArchive(path, base, rar.Deflate)
		return ""
	case tar.Detect(path):
		l.loadArchive(path, base, tar.Deflate)
		return ""
	case zip.Detect(path):
		l.loadArchive(path, base, zip.Deflate)
		return ""
	}

	return path
}

func (l *Loader) process(path string) string {
	//if !flags.CLI.NoPlugins {
	//for _, p := range l.plugins {
	//	if p.Match(path) {
	//		p.Execute(path, func(path, dir string) {
	//			if len(dir) > 0 {
	//				l.loadDir(dir) // load dir results
	//			}
	//
	//			l.addPlugin(path, p.Name)
	//		})
	//
	//		return ""
	//	}
	//}
	//}

	if !flags.CLI.NoConvert {
		if evtx.Detect(path) {
			path = evtx.Parse(path)
		}

		if journal.Detect(path) {
			path = journal.Parse(path)
		}
	}

	return path
}
