package adapter

import (
	"os"
	"path/filepath"
	"time"

	"github.com/cuhsat/fox/internal/opt"
)

type Adapter interface {
	Select(node Node)
	List(root string) []Node
	Width() int
}

type Node struct {
	Leaf  bool
	Data  string
	Text  string
	Value string
}

type FileSystem struct {
	state *opt.State
	open  func(string)
}

func NewFileSystem(state *opt.State, fn func(string)) *FileSystem {
	return &FileSystem{
		state: state,
		open:  fn,
	}
}

func (fs *FileSystem) Select(node Node) {
	fs.state.Call(func() { fs.open(node.Value) })
}

func (fs *FileSystem) List(root string) []Node {
	fs.state.ChangePath(root)

	dir := filepath.Dir(root)

	nodes := []Node{
		{
			Leaf:  false,
			Data:  fs.data(root),
			Text:  ".",
			Value: root,
		},
		{
			Leaf:  false,
			Data:  fs.data(dir),
			Text:  "..",
			Value: dir,
		},
	}

	files, err := os.ReadDir(root)

	if err != nil {
		return nodes
	}

	for _, file := range files {
		path := filepath.Join(root, file.Name())

		nodes = append(nodes, Node{
			Leaf:  !file.IsDir(),
			Data:  fs.data(path),
			Text:  file.Name(),
			Value: path,
		})
	}

	return nodes
}

func (fs *FileSystem) Width() int {
	return 30 // RFC1123
}

func (fs *FileSystem) data(path string) string {
	fi, err := os.Stat(path)

	if err != nil {
		return "error"
	}

	return fi.ModTime().UTC().Format(time.RFC1123)
}
