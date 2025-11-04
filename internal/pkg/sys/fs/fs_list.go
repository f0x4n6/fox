package fs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cuhsat/fox/v3/internal/pkg/text"
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

		dir, err := isDir(path)

		if err != nil {
			continue
		}

		items = append(items, Item{
			Leaf:  !dir,
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
		return fmt.Sprintf(infoFmt, "", "ERR")
	}

	s := text.Human(fi.Size())

	if fi.IsDir() {
		s = "DIR"
	}

	t := fi.ModTime().UTC().Format("02 Jan 15:04")

	return fmt.Sprintf(infoFmt, t, s)
}

func isDir(path string) (bool, error) {
	fi, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeSymlink == 0 {
		return fi.IsDir(), nil
	}

	dst, err := filepath.EvalSymlinks(path)

	if err != nil {
		return false, err
	}

	fi, err = os.Stat(dst)

	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
