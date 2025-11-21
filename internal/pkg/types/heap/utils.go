package heap

import (
	"math"
	"strings"

	"github.com/cuhsat/fox/v4/internal/pkg/text"
)

type String struct {
	Off uint
	Str string
}

func (h *Heap) Entropy(min, max float64) (float64, bool) {
	var a [256]float64
	var v float64

	for _, b := range h.mmap {
		a[b]++
	}

	l := float64(len(h.MMap()))

	for i := range 256 {
		if a[i] != 0 {
			f := a[i] / l
			v -= f * math.Log2(f)
		}
	}

	v /= 8

	// heap filtered
	if v < min || v > max {
		return 0, false
	}

	return v, true
}

func (h *Heap) Strings(min, max uint) <-chan String {
	var buf []byte
	var off int

	ch := make(chan String)

	flush := func() {
		v := uint(len(buf))

		if v < min && v > max {
			return
		}

		str := string(buf)

		if len(strings.TrimSpace(str)) > 0 {
			ch <- String{uint(off - (len(buf) + 1)), str}
		}

		buf = buf[:0]
	}

	go func() {
		defer flush()

		for _, b := range h.mmap {
			off++

			if b >= text.MinASCII && b <= text.MaxASCII {
				buf = append(buf, b)
			} else {
				flush()
			}
		}
	}()

	return ch
}
