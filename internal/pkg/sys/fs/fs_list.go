package fs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cuhsat/fox/internal/pkg/text"
)

const (
	ActualDir = "."
	ParentDir = ".."
)

const InfoWidth = 20

const infoFmt = "%12s %6s"

type Item struct {
	Leaf  bool
	Info  string
	Text  string
	Value string
}

func List(root string) []Item {
	dir := filepath.Dir(root)

	items := []Item{
		{false, Info(root), ActualDir, root},
		{false, Info(dir), ParentDir, dir},
	}

	files, err := os.ReadDir(root)

	if err != nil {
		return items
	}

	for _, file := range files {
		path := filepath.Join(root, file.Name())

		items = append(items, Item{
			Leaf:  !file.IsDir(),
			Info:  Info(path),
			Text:  file.Name(),
			Value: path,
		})
	}

	return items
}

func Info(path string) string {
	fi, err := os.Stat(path)

	if err != nil {
		return fmt.Sprintf(infoFmt, "", "ERROR")
	}

	s := text.Human(fi.Size())

	if fi.IsDir() {
		s = "DIR"
	}

	t := fi.ModTime().UTC().Format("02 Jan 15:04")

	return fmt.Sprintf(infoFmt, t, s)
}
