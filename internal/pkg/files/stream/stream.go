package stream

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

type Stream struct {
	path string
	url  string
	sc   Schema

	f *os.File
}

func New(path, url string, sc Schema) *Stream {
	var err error

	st := Stream{path: path, url: url, sc: sc}

	if len(path) > 0 {
		st.f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			log.Fatal(err)
		}
	}

	return &st
}

func (st *Stream) Close() {
	_ = st.f.Close()
}

func (st *Stream) Write(p []byte) (n int, err error) {
	// stream to file
	if len(st.path) > 0 {
		_, _ = st.f.Write(p)
		Sign(st.path)
	}

	// stream to url
	if len(st.url) > 0 {
		st.sc.Write(string(p))
		Post(st.url, st.sc)
	}

	return len(p), nil
}

func Sign(path string) {
	sha := sha256.New()

	buf, err := os.ReadFile(path)

	if err != nil {
		log.Println(err)
		return
	}

	sha.Write(buf)

	sum := hex.EncodeToString(sha.Sum(nil))

	err = os.WriteFile(path+".sha256", []byte(sum), 0600)

	if err != nil {
		log.Println(err)
	}
}
