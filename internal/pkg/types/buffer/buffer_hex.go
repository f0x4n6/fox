package buffer

import (
	"fmt"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type HexBuffer struct {
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

func Hex(h *heap.Heap, t uint) (buf *HexBuffer) {
	var off uint

	buf = new(HexBuffer)

	mmap := h.MMap()

	if t > 0 {
		off = max(uint(h.Size())-t, 0)
	}

	buf.Lines = make(chan HexLine, Size)

	// stream lines
	go hexStream(buf, off, mmap[:])

	return
}

func hexStream(buf *HexBuffer, o uint, b []byte) {
	defer close(buf.Lines)

	for i := 0; i < len(b); i += 16 {
		nr := fmt.Sprintf("%08x ", o+uint(i))

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

		buf.Lines <- line
	}
}
