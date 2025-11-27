package evtx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"maps"

	"github.com/0xrawsec/golang-evtx/evtx"

	"github.com/cuhsat/fox/v4/internal/pkg/files"
	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
)

const Magic = evtx.EvtxMagic
const Chunk = evtx.ChunkMagic

var hostPath = evtx.Path("/Event/System/Computer")

func Detect(b []byte) bool {
	return files.HasMagic(b, 0, []byte(Magic))
}

func Format(b []byte) ([]byte, error) {
	r, err := evtx.New(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)

	for e := range r.Events() {
		_, err := buf.Write(evtx.ToJSON(e))

		if err != nil {
			log.Println(err)
			continue
		}

		_, err = buf.WriteRune('\n')

		if err != nil {
			log.Println(err)
		}
	}

	_ = r.Close()

	return buf.Bytes(), nil
}

func Decode(rs io.ReadSeeker, off int64) (*evtx.Chunk, error) {
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

func ToEvent(evt *evtx.GoEvtxMap, ext bool) *event.Event {
	var ok bool

	l := event.Event{
		Time:      evt.TimeCreated().UTC(),
		Host:      evt.GetStringStrict(&hostPath),
		Extension: make(map[string]string),
	}

	if l.Message, ok = Events[evt.EventID()]; !ok {
		l.Message = fmt.Sprintf("Undescribed event: Event ID %d", evt.EventID())
	}

	if l.Severity, ok = Levels[evt.EventID()]; !ok {
		l.Severity = 0 // unknown
	}

	if ext {
		for k, v := range maps.All(*evt) {
			l.Extension[k] = fmt.Sprintf("%v", v)
		}
	}

	return &l
}
