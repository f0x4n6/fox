package page

import (
	"fmt"

	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/text"
)

type HexPage struct {
	Page
	Lines chan HexLine
}

type HexLine struct {
	Line
	Hex string
}

func (hl HexLine) String() string {
	// canonical form
	return fmt.Sprintf("%s %s|%-16s|", hl.Nr, hl.Hex, hl.Str)
}

func Hex(ctx *Context) (page *HexPage) {
	var tail int

	page = new(HexPage)

	mmap := *ctx.Heap.MMap()

	limit := flags.Get().Limits

	if limit.IsTail && limit.Bytes > 0 {
		tail = max(int(ctx.Heap.Size())-limit.Bytes, 0)
	}

	page.W, page.H = ctx.W, len(mmap)/16

	if len(mmap)%16 > 0 {
		page.H++
	}

	page.Lines = make(chan HexLine, Size)

	// stream lines
	go hexStream(ctx, page, tail, mmap[ctx.Y*16:])

	return
}

func hexStream(ctx *Context, page *HexPage, t int, b []byte) {
	defer close(page.Lines)

	for i := 0; i < len(b); i += 16 {
		if i/16 >= ctx.H {
			return // page filled
		}

		nr := fmt.Sprintf("%08x ", t+i+(ctx.Y*16))

		line := HexLine{
			Line: Line{Nr: nr, Str: ""},
			Hex:  "",
		}

		for j := range 16 {
			if i+j >= len(b) {
				break
			}

			line.Hex = fmt.Sprintf("%s%02x", line.Hex, b[i+j])
			line.Str = fmt.Sprintf("%s%c", line.Str, b[i+j])

			// make a hex gap every 8 bytes
			if (j+1)%8 == 0 {
				line.Hex += "  "
			} else {
				line.Hex += " "
			}
		}

		line.Hex = fmt.Sprintf("%-*s", 50, line.Hex)
		line.Str = text.ToASCII(line.Str)

		page.Lines <- line
	}
}
