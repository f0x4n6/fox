package plain

import (
	"fmt"
	"log"
	"os"

	"github.com/cuhsat/fox/v3/internal/pkg/files/evidence"
)

const Ext = ".txt"

type Writer struct {
	file *os.File // file handle
}

func New() *Writer {
	return new(Writer)
}

func (w *Writer) Open(f *os.File, _ bool, _ string) {
	w.file = f
}

func (w *Writer) Begin() {}

func (w *Writer) Flush() {}

func (w *Writer) WriteMeta(_ evidence.Meta) {}

func (w *Writer) WriteLine(_, _ int, str string) {
	_, err := fmt.Fprintln(w.file, str)

	if err != nil {
		log.Println(err)
	}
}
