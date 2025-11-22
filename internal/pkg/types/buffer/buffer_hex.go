package buffer

import (
	"fmt"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type HexLine struct {
	Line
	Hex string
}

type HexBuffer struct {
	Lines chan HexLine
}

func (l HexLine) String() string {
	return fmt.Sprintf("%s  %s %-16s", l.Nr, l.Hex, l.Str)
}

func Hex(h *heap.Heap, n uint) *HexBuffer {
	var buf = &HexBuffer{make(chan HexLine, Size)}
	var off uint

	if n > 0 {
		off = max(uint(h.Size())-n, 0)
	}

	go hexStream(buf, off, h.MMap())

	return buf
}

func hexStream(buf *HexBuffer, off uint, b []byte) {
	defer close(buf.Lines)

	for i := 0; i < len(b); i += 16 {
		nr := fmt.Sprintf("%08x ", off+uint(i))

		line := HexLine{Line{Nr: nr, Str: ""}, ""}

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
		line.Str = text.ToAscii(line.Str)

		buf.Lines <- line
	}
}
