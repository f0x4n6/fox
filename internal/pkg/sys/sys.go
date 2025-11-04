package sys

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/rainu/go-command-chain"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
)

const Prefix = "fox:"

var Logs = make(chan string, 64)

type logger struct{}

func (l logger) Write(b []byte) (int, error) {
	msg := strings.TrimSpace(string(b))

	if len(msg) > 0 {
		Logs <- msg
	}

	return len(msg), nil
}

func Init() {
	log.SetFlags(0)
	log.SetOutput(new(logger))
}

func Wait() {
	exit := make(chan os.Signal, 1)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	os.Exit(0)
}

func Exit(v ...any) {
	_, _ = fmt.Fprintf(os.Stderr, Prefix+" %s\n", v...)
	os.Exit(1)
}

func Exec(cmds []string) fs.File {
	f := fs.Create("/fox/exec")
	defer f.Close()

	for _, cmd := range cmds {
		err := cmdchain.Builder().JoinShellCmd(cmd).
			Finalize().WithOutput(f).Run()

		if err != nil {
			log.Println(err)
			break
		}
	}

	return f
}

func Trap() {
	bin, _ := os.Executable()

	args := []string{"/K", bin}
	args = append(args, os.Args[1:]...)

	err := exec.Command("C:\\WINDOWS\\system32\\cmd.exe", args...).Run()

	if err != nil {
		fmt.Printf("%s %s\n", Prefix, err.Error())
	}
}

func Shell() {
	fmt.Println(fox.Product, fox.Version)
	fmt.Println("Type 'exit' to return.")

	shell := os.Getenv("SHELL")

	if len(shell) == 0 {
		if runtime.GOOS == "windows" {
			shell = "C:\\WINDOWS\\system32\\cmd.exe"
		} else {
			shell = "/bin/sh"
		}
	}

	cmd := exec.Command(shell, "-l") // login shell
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
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

func Recover() {
	if err := recover(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, Prefix+" %+v\n\n%s\n", err, debug.Stack())
	}
}
