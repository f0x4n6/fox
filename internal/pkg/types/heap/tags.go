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

func (h *Heap) TagLines(lines string) {
	lines = strings.TrimSpace(lines)
	lines = strings.ToLower(lines)

	switch lines {
	case "a", "all":
		h.TagAll()

	case "n", "none":
		h.TagNone()

	case "i", "invert":
		h.ToggleTags()

	default:
		nrs := h.parse(lines)

		h.Lock()

		for _, nr := range nrs {
			h.tags[nr] = true
		}

		h.Unlock()
	}
}

func (h *Heap) TagAll() {
	fmap := *h.FMap()

	h.Lock()

	for _, str := range fmap {
		h.tags[str.Nr] = true
	}

	h.Unlock()
}

func (h *Heap) TagNone() {
	fmap := *h.FMap()

	h.Lock()

	for _, str := range fmap {
		h.tags[str.Nr] = false
	}

	h.Unlock()
}

func (h *Heap) ToggleTags() {
	fmap := *h.FMap()

	h.Lock()

	for _, str := range fmap {
		h.tags[str.Nr] = !h.tags[str.Nr]
	}

	h.Unlock()
}

func (h *Heap) ToggleTag(nr int) {
	h.Lock()
	h.tags[nr] = !h.tags[nr]
	h.Unlock()
}

func (h *Heap) IsTagged(nr int) bool {
	h.RLock()
	t, ok := h.tags[nr]
	h.RUnlock()

	return ok && t
}
