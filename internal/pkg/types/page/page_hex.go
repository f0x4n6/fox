package page

import (
	"fmt"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type HexPage struct {
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

func Hex(h *heap.Heap) (page *HexPage) {
	var tail int

	page = new(HexPage)

	mmap := h.MMap()

	//limit := flags.CLI.Limits

	//if limit.IsTail && limit.Bytes > 0 {
	//	tail = max(int(h.Size())-limit.Bytes, 0)
	//}

	page.Lines = make(chan HexLine, Size)

	// stream lines
	go hexStream(page, tail, mmap[:])

	return
}

func hexStream(page *HexPage, t int, b []byte) {
	defer close(page.Lines)

	for i := 0; i < len(b); i += 16 {
		nr := fmt.Sprintf("%08x ", t+i)

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
