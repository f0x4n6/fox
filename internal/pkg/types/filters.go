package types

import (
	"regexp"

	"github.com/cuhsat/fox/v4/internal/pkg/types/smap"
)

type Filters struct {
	Regex  *regexp.Regexp // regex pattern
	Before int            // lines before
	After  int            // lines after
}

func (f *Filters) FilterSMap(s smap.SMap) smap.SMap {
	if f.Regex == nil {
		return s // not filtered
	}

	v := s.Grep(f.Regex)

	if f.Before+f.After == 0 {
		return v // without context
	}

	r := make(smap.SMap, len(v))

	for grp, str := range v {
		for _, b := range (s)[max((str.Nr-1)-f.Before, 0) : str.Nr-1] {
			b.Grp = grp + 1
			r = append(r, b)
		}

		str.Grp = grp + 1
		r = append(r, str)

		for _, a := range (s)[str.Nr:min(str.Nr+f.After, len(s))] {
			a.Grp = grp + 1
			r = append(r, a)
		}
	}

	return r // with context
}
