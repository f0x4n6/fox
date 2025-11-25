package hunt

import (
	"bufio"
	"bytes"
	"fmt"
	"hash/maphash"
	"io"
	"log"
	"regexp"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

const Windows = "/Windows/System32/winevt/Logs"

var re = regexp.MustCompile(evtx.ChunkMagic)

func Hunt(h *heap.Heap, v int) chan *Log {
	ch := make(chan *Log, 1024)

	r1 := bytes.NewReader(h.MMap())
	r2 := bytes.NewReader(h.MMap())

	evtx.SetModeCarving(true)

	set := make(types.Set)

	go func() {
		defer close(ch)

		seed := maphash.MakeSeed()

		for off := range offset(r1) {
			if v > 2 {
				log.Printf("parsing chunk 0x%08x\n", off)
			}

			chunk, err := parse(r2, off)

			if err != nil {
				log.Print(err)
				continue
			}

			for evt := range chunk.Events() {
				key := maphash.String(seed, fmt.Sprintf("%#v", evt))

				if _, ok := set[key]; !ok {
					ch <- Transform(evt)
					set[key] = types.Nil
				}
			}
		}
	}()

	return ch
}

func offset(rs io.ReadSeeker) <-chan int64 {
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

func parse(rs io.ReadSeeker, off int64) (*evtx.Chunk, error) {
	evtx.GoToSeeker(rs, off)

	chk := evtx.NewChunk()
	chk.Offset = off
	chk.Data = make([]byte, evtx.ChunkSize)

	if _, err := rs.Read(chk.Data); err != nil {
		return nil, err
	}

	r := bytes.NewReader(chk.Data)

	chk.ParseChunkHeader(r)

	if err := chk.Header.Validate(); err != nil {
		return nil, err
	}

	evtx.GoToSeeker(r, int64(chk.Header.SizeHeader))

	chk.ParseStringTable(r)

	if err := chk.ParseTemplateTable(r); err != nil {
		return nil, err
	}

	if err := chk.ParseEventOffsets(r); err != nil {
		return nil, err
	}

	return &chk, nil
}
