package sys

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"

	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
)

const Prefix = "fox:"

func Exit(v ...any) {
	_, _ = fmt.Fprintf(os.Stderr, Prefix+" %s\n", v...)
	os.Exit(1)
}

func Stdin() fs.File {
	if !Piped(os.Stdin) {
		log.Panicln("Device mode is invalid")
	}

	f := fs.Create("/fox/stdin")

	go func(f fs.File) {
		r := bufio.NewReader(os.Stdin)

		for {
			s, err := r.ReadString('\n')

			switch err {
			case nil:
				_, err = f.WriteString(s)

				if err != nil {
					log.Println(err)
				}

			case io.EOF:
				_ = f.Close()
				break

			default:
				log.Println(err)
			}
		}
	}(f)

	return f
}

func Piped(file fs.File) bool {
	fi, err := file.Stat()

	if err != nil {
		log.Println(err)
		return false
	}

	return (fi.Mode() & os.ModeCharDevice) != os.ModeCharDevice
}

func Handler(fn func() error) {
	_ = fn()
}

func Recover() {
	if err := recover(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, Prefix+" %+v\n\n%s\n", err, debug.Stack())
	}
}
