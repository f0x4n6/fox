package page

import (
	"fmt"

	"github.com/cuhsat/fox/v3/internal/pkg/text"
	"github.com/cuhsat/fox/v3/internal/pkg/types/smap"
)

type TextPage struct {
	Page
	Y     int
	N     int
	FMap  *smap.SMap
	Lines chan *TextLine
}

type TextLine struct {
	Line
	Parts []Part
}

type entry struct {
	s *smap.SMap
	w int // width
	h int // height
}

func (tl TextLine) String() string {
	return tl.Str
}

func Text(ctx *Context) (page *TextPage) {
	page = new(TextPage)

	page.Lines = make(chan *TextLine, Size)
	page.N = text.Dec(ctx.Heap.Length())
	page.Y = ctx.Y

	// recalculate render area
	if ctx.Navi {
		ctx.W -= 2 + page.N
		ctx.H -= 1
	}

	key := ctx.Hash()

	// cache transformed smap to improve performance
	if val, ok := ctx.Heap.Cache.Load(key); ok {
		page.FMap = val.(entry).s
		page.W = val.(entry).w
		page.H = val.(entry).h
	} else {
		page.FMap = ctx.Heap.FMap()

		if ctx.Wrap && page.FMap.CanFormat() {
			page.FMap = page.FMap.Format(ctx.Space)
		} else if ctx.Wrap {
			page.FMap = page.FMap.Wrap(ctx.Space, ctx.W)
		} else {
			page.FMap = page.FMap.Render(ctx.Space)
		}

		page.W, page.H = page.FMap.Size()

		ctx.Heap.Cache.Store(key, entry{
			page.FMap,
			page.W,
			page.H,
		})
	}

	// restore last position
	if ctx.Nr > 0 {
		lastY := max(len(*page.FMap)-1, 0)

		// find the requested line
		page.Y, _ = page.FMap.Find(ctx.Nr)
		page.Y = min(page.Y, lastY)
	}

	go textStream(ctx, page)

	return
}

func textLine(nr, str string, grp int, tag bool) *TextLine {
	return &TextLine{
		Line{nr, grp, tag, str},
		make([]Part, 0),
	}
}

func textStream(ctx *Context, page *TextPage) {
	defer close(page.Lines)

	numSep, numGrp, lastGrp := 0, 1, 0

	// pin first line
	if ctx.Sticky {
		nr := fmt.Sprintf("%0*d ", page.N, 0)
		str := (*ctx.Heap.SMap())[0].Str
		str = text.Trim(str, min(ctx.X, text.Len(str)), ctx.W)

		page.Lines <- textLine(nr, str, 0, false)
	}

	// stream lines
	for y, str := range (*page.FMap)[page.Y:] {
		if y >= ctx.H {
			return // page filled
		}

		// insert context separator
		if ctx.Context && lastGrp != str.Grp && numGrp > 1 {
			page.Lines <- textLine(Sep, "", str.Grp, false)
			numGrp = 1
			numSep++
		}

		off := min(ctx.X, text.Len(str.Str))

		// build line
		line := textLine(
			fmt.Sprintf("%0*d ", page.N, str.Nr),
			text.Trim(str.Str, off, ctx.W),
			str.Grp,
			ctx.Heap.IsTagged(str.Nr),
		)

		if ctx.Sticky {
			y++ // skip first line
		}

		// build parts
		for _, f := range ctx.Heap.Filters() {
			if f.Pattern.Regex == nil {
				continue // skip picked lines
			}

			// find parts
			for _, i := range f.Pattern.Regex.FindAllStringIndex(str.Str, -1) {
				line.Parts = append(line.Parts, Part{
					text.Len(str.Str[:i[0]]) - off,
					y + numSep,
					str.Grp,
					str.Str[i[0]:i[1]],
				})
			}
		}

		page.Lines <- line

		lastGrp = str.Grp
		numGrp++
	}
}
