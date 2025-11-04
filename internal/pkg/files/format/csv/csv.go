package csv

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/ntauth/better-csvd"

	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v3/internal/pkg/text"
)

func Detect(path string) bool {
	ext := filepath.Ext(strings.ToLower(path))
	return ext == ".csv" || ext == ".tsv"
}

func Format(path string) string {
	f := fs.Open(path)
	defer f.Close()

	t := fs.Create(path)
	defer t.Close()

	r := csvd.NewReader(f)
	r.LazyQuotes = true

	cols, err := r.ReadAll()

	if err != nil {
		log.Println(err)
		return path
	}

	ls := make([]int, 0)

	// calculate row max length
	for _, rows := range cols {
		for i, row := range rows {
			if len(ls) < i+1 {
				ls = append(ls, 0)
			}

			ls[i] = max(text.Len(row), ls[i])
		}
	}

	var sb strings.Builder

	// pad all rows
	for _, rows := range cols {
		for i, row := range rows {
			sb.WriteString(text.Pad(row, ls[i]+1))
		}

		sb.WriteRune('\n')

		_, err := t.WriteString(sb.String())

		if err != nil {
			log.Println(err)
		}

		sb.Reset()
	}

	return t.Name()
}
