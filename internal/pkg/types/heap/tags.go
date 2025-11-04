package heap

import (
	"strings"
)

type Tags map[int]bool

func (h *Heap) HasTags() bool {
	h.RLock()
	defer h.RUnlock()

	for _, str := range *h.smap {
		if h.tags[str.Nr] {
			return true
		}
	}

	return false
}

func (h *Heap) IsTagged(nr int) bool {
	h.RLock()
	t, ok := h.tags[nr]
	h.RUnlock()

	return ok && t
}

func (h *Heap) TagLine(nr int, v bool) {
	h.Lock()
	h.tags[nr] = v
	h.Unlock()
}

func (h *Heap) TagLines(lines string) {
	lines = strings.TrimSpace(lines)
	lines = strings.ToLower(lines)

	switch lines {
	case "a", "all":
		h.TagAll()

	case "n", "none":
		h.TagNone()

	case "i", "invert":
		h.TagOthers()

	default:
		nrs := h.parse(lines)

		for _, nr := range nrs {
			h.TagLine(nr, true)
		}
	}
}

func (h *Heap) TagAll() {
	fmap := *h.FMap()

	for _, str := range fmap {
		h.TagLine(str.Nr, true)
	}
}

func (h *Heap) TagNone() {
	fmap := *h.FMap()

	for _, str := range fmap {
		h.TagLine(str.Nr, false)
	}
}

func (h *Heap) TagOthers() {
	fmap := *h.FMap()

	h.Lock()

	for _, str := range fmap {
		h.tags[str.Nr] = !h.tags[str.Nr]
	}

	h.Unlock()
}
