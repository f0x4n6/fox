package page

import (
	"fmt"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v4/internal/pkg/types/smap"
)

type TextPage struct {
	Y     int
	N     int
	SMap  smap.SMap
	Lines chan *TextLine
}

type TextLine struct {
	Line
	Org int
}

func (tl TextLine) String() string {
	return tl.Str
}

func Text(h *heap.Heap, e int) (page *TextPage) {
	page = new(TextPage)

	page.Lines = make(chan *TextLine, Size)
	page.N = text.Dec(h.Len())

	page.SMap = h.SMap()

	if page.SMap.CanFormat() {
		page.SMap = page.SMap.Format(e)
	} else {
		page.SMap = page.SMap.Render(e)
	}

	go textStream(page)

	return
}

func textLine(nr, str string, org, grp int) *TextLine {
	return &TextLine{
		Line{nr, grp, str},
		org,
	}
}

func textStream(page *TextPage) {
	defer close(page.Lines)

	numSep, numGrp, lastGrp := 0, 1, 0

	// stream lines
	for _, str := range page.SMap[page.Y:] {

		// insert context separator
		if lastGrp != str.Grp && numGrp > 1 {
			page.Lines <- textLine(Sep, "", str.Nr, str.Grp)
			numGrp = 1
			numSep++
		}

		// build line
		line := textLine(
			fmt.Sprintf("%0*d ", page.N, str.Nr),
			str.Str,
			str.Nr,
			str.Grp,
		)

		page.Lines <- line

		lastGrp = str.Grp
		numGrp++
	}
}
