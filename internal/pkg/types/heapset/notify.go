package heapset

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
)

func (hs *HeapSet) SetChanged(fn Changed) {
	hs.changed = fn
}

func (hs *HeapSet) WatchFiles() {
	hs.Range(func(_ int, h *heap.Heap) bool {
		hs.watchFile(h.Path)
		return true
	})
}

func (hs *HeapSet) watchFile(path string) {
	err := fs.Watcher.Add(filepath.Dir(path))

	if err != nil {
		log.Println(err)
	}
}

func (hs *HeapSet) notify() {
	for {
		select {
		case ev, ok := <-fs.Watcher.Events:
			if !ok || !ev.Has(fsnotify.Write) {
				continue
			}

			if idx, ok := hs.findByPath(ev.Name); ok {
				h := hs.atomicGet(idx)
				h.Reload()

				if hs.changed != nil {
					hs.changed(h) // raise changed
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
