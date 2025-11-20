package sys

import (
	"log"
	"runtime/debug"
)

const Prefix = "fox: "

func Ignore(fn func() error) {
	_ = fn()
}

func Recover() {
	if err := recover(); err != nil {
		log.Printf("%+v\n\n%s\n", err, debug.Stack())
	}
}
