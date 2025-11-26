package evtx

import (
	"bytes"
	"io"
	"regexp"

	"github.com/0xrawsec/golang-evtx/evtx"
)

var Magic = regexp.MustCompile(evtx.ChunkMagic)

func Parse(rs io.ReadSeeker, off int64) (*evtx.Chunk, error) {
	evtx.SetModeCarving(true)
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
