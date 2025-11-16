package evidence

import (
	"os"
	"os/user"
	"time"
)

type Writer interface {
	Open(file *os.File, old bool, title string)

	Begin()
	Flush()

	WriteMeta(meta Meta)
	WriteLine(nr, grp int, str string)
}

type Meta struct {
	User     *user.User `json:"user"`
	Path     string     `json:"path"`
	Size     int64      `json:"size"`
	Hash     []byte     `json:"hash"`
	Modified time.Time  `json:"modified"`
	Seized   time.Time  `json:"seized"`
}

type Line struct {
	Nr  int    `json:"nr"`
	Grp int    `json:"grp"`
	Str string `json:"str"`
}

func Utc(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
