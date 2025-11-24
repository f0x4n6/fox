package hunt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"

	"github.com/0xrawsec/golang-evtx/evtx"
	"github.com/zeebo/xxh3"

	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

var re = regexp.MustCompile(evtx.ChunkMagic)

func Hunt(h *heap.Heap, v int) chan string {
	ch := make(chan string, 1024)

	r1 := bytes.NewReader(h.MMap())
	r2 := bytes.NewReader(h.MMap())

	evtx.SetModeCarving(true)

	cache := make(Logs)

	go func() {
		defer close(ch)

		for o := range offset(r1) {
			if v > 2 {
				log.Printf("Parsing chunk 0x%08x\n", o)
			}

			chunk, err := parse(r2, o)

			if err != nil {
				log.Print(err)
				continue
			}

			for e := range chunk.Events() {
				s := fmt.Sprintf("%#v", e)

				key := xxh3.HashString(s)

				if _, ok := cache[key]; !ok {
					p := evtx.Path("/Event/System/Computer")

					user, _ := e.UserID()
					host, _ := e.GetString(&p)

					msg, ok := DB[e.EventID()]

					if !ok {
						msg = "not found"
					}

					evt := &Log{
						Ts:      e.TimeCreated(),
						Id:      e.EventID(),
						User:    user,
						Host:    host,
						Channel: e.Channel(),
						Message: msg,
					}

					cache[key] = evt

					ch <- evt.String()
				}
			}
		}

		fmt.Printf("found %d events\n", len(cache))
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
