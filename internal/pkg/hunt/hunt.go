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

	"github.com/cuhsat/fox/v4/internal/pkg/files/parser/evtx"
	"github.com/cuhsat/fox/v4/internal/pkg/files/parser/journal"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

var Paths = []string{
	"/Windows/System32/winevt/Logs",
	"/var/log/journal",
	"/run/log/journal",
}

func Hunt(h *heap.Heap, v int) chan *Log {
	var wg sync.WaitGroup

	ch := make(chan *Log, 1024)

	wg.Add(2)

	go func() {
		defer wg.Done()

		seed := maphash.MakeSeed()

		r1 := bytes.NewReader(h.MMap())
		r2 := bytes.NewReader(h.MMap())

		set := make(types.Set)

		for off := range offset(r1, evtx.Magic) {
			if v > 2 {
				log.Printf("parsing offset 0x%08x\n", off)
			}

			chunk, err := evtx.Parse(r2, off)

			if err != nil {
				log.Print(err)
				continue
			}

			for evt := range chunk.Events() {
				key := maphash.String(seed, fmt.Sprintf("%#v", evt))

				if _, ok := set[key]; !ok {
					ch <- FromEvtx(evt)
					set[key] = types.Nil
				}
			}
		}
	}()

	go func() {
		defer wg.Done()

		seed := maphash.MakeSeed()

		r := bytes.NewReader(h.MMap())

		set := make(types.Set)

		for off := range offset(r, journal.Magic) {
			if v > 2 {
				log.Printf("parsing offset 0x%08x\n", off)
			}

			jfile, err := journal.Parse(h.MMap(), off)

			if err != nil {
				log.Print(err)
				continue
			}

			for jlg := range jfile.GetLogs(context.Background()) {
				key := maphash.String(seed, fmt.Sprintf("%#v", jlg))

				fmt.Printf("%#v\n", jlg)

				if _, ok := set[key]; !ok {
					ch <- FromJournal(jlg)
					set[key] = types.Nil
				}
			}
		}
	}()

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}

func offset(rs io.ReadSeeker, re *regexp.Regexp) <-chan int64 {
	ch := make(chan int64, 1024)

	go func(r *bufio.Reader) {
		var lst int64

		for loc := re.FindReaderIndex(r); loc != nil; loc = re.FindReaderIndex(r) {
			off, _ := rs.Seek(0, io.SeekCurrent)
			ch <- lst + int64(loc[0])
			lst = off - int64(r.Buffered())
		}

		close(ch)
	}(bufio.NewReader(rs))

	return ch
}
