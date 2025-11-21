package buffer

import (
	"fmt"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/types/smap"
)

type TextBuffer struct {
	Y     int
	N     int
	SMap  smap.SMap
	Lines chan *TextLine
}

type TextLine struct {
	Line
}

func (tl TextLine) String() string {
	return tl.Str
}

func Text(h *heap.Heap, e int) (buf *TextBuffer) {
	buf = new(TextBuffer)

	buf.Lines = make(chan *TextLine, Size)
	buf.N = text.Dec(h.Len())

	buf.SMap = h.SMap()

	if buf.SMap.CanFormat() {
		buf.SMap = buf.SMap.Format(e)
	} else {
		buf.SMap = buf.SMap.Render(e)
	}

	go textStream(buf)

	return
}

func textLine(nr, str string, grp uint) *TextLine {
	return &TextLine{Line{nr, grp, str}}
}

func textStream(buf *TextBuffer) {
	defer close(buf.Lines)

	var numSep uint = 0
	var numGrp uint = 1
	var tmpGrp uint = 0

	// stream lines
	for _, str := range buf.SMap[buf.Y:] {

		// insert context separator
		if tmpGrp != str.Grp && numGrp > 1 {
			buf.Lines <- textLine(Sep, "", str.Grp)
			numGrp = 1
			numSep++
		}

		// build line
		line := textLine(
			fmt.Sprintf("%0*d ", buf.N, str.Nr),
			str.Str,
			str.Grp,
		)

		buf.Lines <- line

		tmpGrp = str.Grp
		numGrp++
	}
}
