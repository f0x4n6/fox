package sys

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

const Prefix = "fox:"

var Logs = make(chan string, 64)

func Setup() {
	log.SetFlags(0)
	log.SetOutput(new(logger))
}

type logger struct{}

func (l logger) Write(b []byte) (int, error) {
	msg := strings.TrimSpace(string(b))

	if len(msg) > 0 {
		Logs <- msg
	}

	return 0, nil
}

func Trace() {
	if err := recover(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf(Prefix+" %+v\n\n%s", err, debug.Stack()))
	}
}

func Exit(v ...any) {
	_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf(Prefix+" %s", v...))
	os.Exit(1)
}

func Error(v ...any) {
	log.Println(v...)
}
