package sys

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
)

const Prefix = "fox: "

func Stdin() ([]byte, error) {
	if !Piped(os.Stdin) {
		return nil, errors.New("invalid device mode")
	}

	b, err := io.ReadAll(bufio.NewReader(os.Stdin))

	if err != nil {
		return nil, err
	}

	return b, nil
}

func Piped(file *os.File) bool {
	fi, err := file.Stat()

	if err != nil {
		log.Println(err)
		return false
	}

	return (fi.Mode() & os.ModeCharDevice) != os.ModeCharDevice
}

func Handle(fn func() error) {
	_ = fn()
}

func Recover() {
	if err := recover(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, Prefix+"%+v\n\n%s\n", err, debug.Stack())
	}
}
