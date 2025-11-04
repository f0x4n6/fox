package heap

import (
	"math"
	"regexp"
	"strings"

	"github.com/cuhsat/fox/v3/internal/pkg/text"
)

func (h *Heap) Entropy(n, m float64) float64 {
	var a [256]float64
	var v float64

	ch := make(chan byte, 1024)

	go h.stream(ch)

	for b := range ch {
		a[b]++
	}

	l := float64(len(*h.MMap()))

	for i := range 256 {
		if a[i] != 0 {
			f := a[i] / l
			v -= f * math.Log2(f)
		}
	}

	v /= 8

	if v < n || v > m {
		return -1 // filtered
	}

	return v
}

func (h *Heap) Strings(n, m int, c bool, re *regexp.Regexp) <-chan text.String {
	ch := make(chan byte, 1024)

	str := make(chan text.String)
	cls := make(chan text.String)

	go h.stream(ch)
	go text.Carve(ch, str, n, m)

	if !c && re == nil {
		return str
	}

	go text.Match(str, cls, c, re)

	return cls
}

func (h *Heap) stream(ch chan<- byte) {
	h.RLock()

	for _, b := range *h.mmap {
		ch <- b
	}

	h.RUnlock()

	close(ch)
}

func (h *Heap) parse(lines string) (nrs []int) {
	for _, l := range strings.Split(lines, ",") {
		r := strings.Split(l, "-")

		if strings.HasPrefix(l, "%") {
			n := text.Int(strings.TrimPrefix(l, "%"))

			for i := n; i <= h.Length(); i += n {
				nrs = append(nrs, i)
			}
		} else if len(r) > 1 {
			a := text.Int(r[0])
			b := text.Int(r[1])

			if a > 0 && b > 0 && a <= b {
				for i := a; i <= b; i++ {
					nrs = append(nrs, i)
				}
			}
		} else {
			if nr := text.Int(l); nr > 0 {
				nrs = append(nrs, nr)
			}
		}
	}

	return
}
