package heapset

import (
	"log"
	"path/filepath"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"

	"github.com/cuhsat/fox/internal/pkg/sys/fs"
)

func (hs *HeapSet) SetCallback(fn Callback) {
	hs.watch = fn
}

func (hs *HeapSet) addFile(path string) {
	err := fs.Watcher.Add(filepath.Dir(path))

	if err != nil {
		log.Println(err)
	}
}

func (hs *HeapSet) watchFiles() {
	for {
		select {
		case ev, ok := <-fs.Watcher.Events:
			if !ok || !ev.Has(fsnotify.Write) {
				continue
			}

			idx, ok := hs.findByPath(ev.Name)

			if ok && idx == atomic.LoadInt32(hs.index) {
				h := hs.atomicGet(idx)
				h.Reload()

				if hs.watch != nil {
					hs.watch() // raise watch
				}

				continue
			}

		case err, ok := <-fs.Watcher.Errors:
			if ok {
				log.Println(err)
			}
		}
	}
}
