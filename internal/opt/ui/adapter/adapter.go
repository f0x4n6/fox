package adapter

import (
	"os"
	"path/filepath"

	"github.com/cuhsat/fox/internal/opt"
)

type Callback func(string)

type Adapter interface {
	List(root string) []Node
}

type Node struct {
	Leaf  bool
	Text  string
	Value string
}

func (n *Node) String() string {
	return n.Text
}

type FileSystem struct {
	state *opt.State
}

func New(state *opt.State) *FileSystem {
	return &FileSystem{
		state: state,
	}
}

func (fs *FileSystem) List(root string) []Node {
	fs.state.ChangePath(root)

	nodes := []Node{
		{
			Leaf:  false,
			Text:  ".",
			Value: root,
		},
		{
			Leaf:  false,
			Text:  "..",
			Value: filepath.Dir(root),
		},
	}

	files, err := os.ReadDir(root)

	if err != nil {
		return nodes
	}

	for _, file := range files {
		nodes = append(nodes, Node{
			Leaf:  !file.IsDir(),
			Text:  file.Name(),
			Value: filepath.Join(root, file.Name()),
		})
	}

	return nodes
}
