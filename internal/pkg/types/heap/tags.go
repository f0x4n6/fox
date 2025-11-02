package heap

type Tags map[int]bool

func (h *Heap) TagAll() {
	for _, str := range *h.smap {
		h.tags[str.Nr] = true
	}
}

func (h *Heap) TagNone() {
	for _, str := range *h.smap {
		h.tags[str.Nr] = false
	}
}

func (h *Heap) InvertTags() {
	for _, str := range *h.smap {
		h.tags[str.Nr] = !h.tags[str.Nr]
	}
}

func (h *Heap) InvertTag(nr int) {
	h.tags[nr] = !h.tags[nr]
}

func (h *Heap) IsTagged(nr int) bool {
	t, ok := h.tags[nr]

	return ok && t
}
