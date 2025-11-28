package hunt

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"regexp"
	"sync"

	"github.com/cuhsat/fox/v4/internal/pkg/files/format/evtx"
	"github.com/cuhsat/fox/v4/internal/pkg/files/format/journal"
	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

const (
	Level = 8
	size  = 1024
)

var Paths = []string{
	"/Windows/System32/winevt/Logs",
	"/var/log/journal",
	"/run/log/journal",
}

type Context struct {
	h *heap.Heap // heap
	x int        // cli ext
	v int        // cli verbose
}

func Hunt(h *heap.Heap, x int, v int) chan *event.Event {
	ch := make(chan *event.Event, size)
	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx := &Context{h, x, v}

	// hunt Windows Event Logs
	go func() {
		defer wg.Done()

		r1 := bytes.NewReader(h.MMap())
		r2 := bytes.NewReader(h.MMap())

		re := regexp.MustCompile(evtx.Chunk)

		for off := range offset(ctx, r1, re) {
			for evt := range evtx.Decode(r2, off, x) {
				ch <- evt
			}
		}
	}()

	// hunt Linux Systemd journals
	go func() {
		defer wg.Done()

		r := bytes.NewReader(h.MMap())

		re := regexp.MustCompile(journal.Magic)

		for off := range offset(ctx, r, re) {
			for evt := range journal.Decode(h.MMap(), off, x) {
				ch <- evt
			}
		}
	}()

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}

func offset(ctx *Context, rs io.ReadSeeker, re *regexp.Regexp) <-chan int64 {
	ch := make(chan int64, size)

	go func(r *bufio.Reader) {
		var lst int64

		for loc := re.FindReaderIndex(r); loc != nil; loc = re.FindReaderIndex(r) {
			off, _ := rs.Seek(0, io.SeekCurrent)
			ch <- lst + int64(loc[0])
			lst = off - int64(r.Buffered())

			if ctx.v > 2 {
				log.Printf("parsing offset 0x%08x\n", loc[1])
			}
		}

		close(ch)
	}(bufio.NewReader(rs))

	return ch
}
