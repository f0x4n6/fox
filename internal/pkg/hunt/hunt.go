package hunt

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"hash/maphash"
	"io"
	"log"
	"regexp"
	"sync"

	"github.com/cuhsat/fox/v4/internal/pkg/files/format/evtx"
	"github.com/cuhsat/fox/v4/internal/pkg/files/format/journal"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

const (
	Limit = 8
	size  = 1024
)

var Paths = []string{
	"/Windows/System32/winevt/Logs",
	"/var/log/journal",
	"/run/log/journal",
}

type Context struct {
	h  *heap.Heap
	x  bool
	v  int
	ch chan<- *event.Event
	wg *sync.WaitGroup
}

func Hunt(h *heap.Heap, x bool, v int) chan *event.Event {
	ch := make(chan *event.Event, size)

	ctx := &Context{h, x, v, ch, new(sync.WaitGroup)}
	ctx.wg.Add(2)

	go huntEventlogs(ctx)
	go huntJournals(ctx)

	go func() {
		defer close(ch)
		ctx.wg.Wait()
	}()

	return ch
}

func huntEventlogs(ctx *Context) {
	defer ctx.wg.Done()

	seed := maphash.MakeSeed()

	r1 := bytes.NewReader(ctx.h.MMap())
	r2 := bytes.NewReader(ctx.h.MMap())

	set := make(types.Set)

	re := regexp.MustCompile(evtx.Chunk)

	for off := range offset(ctx, r1, re) {
		chunk, err := evtx.Decode(r2, off)

		if err != nil {
			log.Print(err)
			continue
		}

		for evt := range chunk.Events() {
			key := maphash.String(seed, fmt.Sprintf("%#v", evt))

			if _, ok := set[key]; !ok {
				ctx.ch <- evtx.ToEvent(evt, ctx.x)
				set[key] = types.Nil
			}
		}
	}
}

func huntJournals(ctx *Context) {
	defer ctx.wg.Done()

	seed := maphash.MakeSeed()

	r := bytes.NewReader(ctx.h.MMap())

	set := make(types.Set)

	re := regexp.MustCompile(journal.Magic)

	for off := range offset(ctx, r, re) {
		jfile, err := journal.Decode(ctx.h.MMap(), off)

		if err != nil {
			log.Print(err)
			continue
		}

		for jlg := range jfile.GetLogs(context.Background()) {
			key := maphash.String(seed, fmt.Sprintf("%#v", jlg))

			if _, ok := set[key]; !ok {
				ctx.ch <- journal.ToEvent(jlg, ctx.x)
				set[key] = types.Nil
			}
		}
	}
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
