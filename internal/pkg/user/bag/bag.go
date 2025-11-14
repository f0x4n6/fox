package bag

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/cuhsat/fox/v4/internal"
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence"
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence/json"
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence/sqlite"
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence/text"
	"github.com/cuhsat/fox/v4/internal/pkg/files/evidence/url"
	"github.com/cuhsat/fox/v4/internal/pkg/files/schema/ecs"
	"github.com/cuhsat/fox/v4/internal/pkg/files/schema/hec"
	"github.com/cuhsat/fox/v4/internal/pkg/files/schema/raw"
	"github.com/cuhsat/fox/v4/internal/pkg/flags"
	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type Bag struct {
	Path string // file path
	Mode string // file mode

	file *os.File          // file handle
	name string            // case name
	key  string            // key phrase
	url  string            // url address
	ws   []evidence.Writer // writers
}

func New() *Bag {
	var ws []evidence.Writer

	cli := &flags.CLI

	path := cli.File

	if len(path) == 0 {
		path = cli.File
	}

	switch cli.Mode {
	case types.SQLITE:
		ws = append(ws, sqlite.New())
		path += sqlite.Ext

	case types.JSONL:
		ws = append(ws, json.New(false))
		path += json.ExtPretty

	case types.JSON:
		ws = append(ws, json.New(true))
		path += json.Ext

	case types.TEXT:
		ws = append(ws, text.New())
		path += text.Ext

	default:
		// write nothing
	}

	if len(cli.Url) > 0 {
		if cli.Ecs {
			ws = append(ws, url.New(cli.Url, ecs.New()))
		} else if cli.Hec {
			ws = append(ws, url.New(cli.Url, hec.New()))
		} else {
			ws = append(ws, url.New(cli.Url, raw.New()))
		}
	}

	return &Bag{
		Path: path,
		Mode: cli.Mode,
		name: cli.Case,
		key:  cli.Sign,
		url:  cli.Url,
		file: nil,
		ws:   ws,
	}
}

func (bag *Bag) String() string {
	if bag.file != nil {
		return bag.Path
	} else {
		return bag.url
	}
}

func (bag *Bag) Close() {
	if bag.file != nil {
		_ = bag.file.Close()
	}
}

func (bag *Bag) Put(h *heap.Heap) bool {
	bag.init()

	usr, err := user.Current()

	if err != nil {
		log.Println(err)
	}

	sum, err := h.HashSum(types.SHA256)

	if err != nil {
		log.Println(err)
	}

	abs, err := filepath.Abs(h.String())

	if err != nil {
		log.Println(err)
	}

	var ptn []string

	for _, p := range h.Patterns() {
		ptn = append(ptn, p.Value)
	}

	for _, w := range bag.ws {
		w.Begin()

		w.WriteMeta(evidence.Meta{
			User:     usr,
			Name:     bag.name,
			Path:     abs,
			Size:     h.Size(),
			Hash:     sum,
			Filters:  ptn,
			Seized:   now(),
			Modified: mod(h),
		})

		for _, str := range *h.SMap() {
			w.WriteLine(str.Nr, str.Grp, str.Str)
		}

		w.Flush()
	}

	if bag.file != nil {
		bag.sign()
	}

	return len(bag.ws) > 0
}

func (bag *Bag) init() {
	old := fs.Exists(bag.Path)

	if bag.Mode != types.NONE {
		var err error

		bag.file, err = os.OpenFile(bag.Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)

		if err != nil {
			log.Panicln(err)
		}
	}

	title := fmt.Sprintf("Forensic Examiner Evidence Bag (%s)", fox.Version)

	for _, w := range bag.ws {
		w.Open(bag.file, old, title)
	}
}

func (bag *Bag) sign() {
	var imp hash.Hash

	if len(bag.key) > 0 {
		imp = hmac.New(sha256.New, []byte(bag.key))
	} else {
		imp = sha256.New()
	}

	buf, err := os.ReadFile(bag.Path)

	if err != nil {
		log.Println(err)
		return
	}

	imp.Write(buf)

	sum := base64.StdEncoding.EncodeToString(imp.Sum(nil))

	err = os.WriteFile(bag.Path+".sig", []byte(sum), 0600)

	if err != nil {
		log.Println(err)
	}

	return
}

func now() time.Time {
	return time.Now().UTC()
}

func mod(h *heap.Heap) time.Time {
	mt := now()

	if h.Type == types.Regular {
		fi, err := os.Stat(h.Base)

		if err == nil {
			mt = fi.ModTime()
		} else {
			log.Println(err)
		}
	}

	return mt
}
