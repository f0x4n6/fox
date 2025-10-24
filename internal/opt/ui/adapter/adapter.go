package adapter

import (
	"log"
	"os"
	"path/filepath"
)

type Adapter interface {
	Init() []Node
	List(root string) []Node
}

type Callback func(string)

type Node struct {
	Leaf  bool
	Text  string
	Value string
}

type Filesystem struct {
}

func (fs *Filesystem) Init() []Node {
	dir, err := os.Getwd()

	if err != nil {
		log.Panicln(err)
	}

	return fs.List(dir)
}

func (fs *Filesystem) List(root string) []Node {
	nodes := []Node{
		{
			Leaf:  true,
			Text:  ".",
			Value: root,
		},
		{
			Leaf:  true,
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
