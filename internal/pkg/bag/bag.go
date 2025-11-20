package bag

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
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
	"github.com/cuhsat/fox/v4/internal/pkg/types"
	"github.com/cuhsat/fox/v4/internal/pkg/types/heap"
)

type Options struct {
	File string
	Mode string
	Sign string
	Auth string
	Url  string
	ECS  bool
	HEC  bool
}

type Bag struct {
	opts *Options          // bag options
	path string            // file path
	file *os.File          // file handle
	ws   []evidence.Writer // writers
}

func New(opts *Options) *Bag {
	var ws []evidence.Writer

	path := opts.File

	switch opts.Mode {
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

	if len(opts.Url) > 0 {
		if opts.ECS {
			ws = append(ws, url.New(opts.Url, ecs.New()))
		} else if opts.HEC {
			ws = append(ws, url.New(opts.Url, hec.New(opts.Auth)))
		} else {
			ws = append(ws, url.New(opts.Url, raw.New()))
		}
	}

	return &Bag{
		opts: opts,
		path: path,
		file: nil,
		ws:   ws,
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

	for _, w := range bag.ws {
		w.Begin()

		w.WriteMeta(evidence.Meta{
			User:     usr,
			Path:     abs,
			Size:     h.Size(),
			Hash:     sum,
			Seized:   time.Now().UTC(),
			Modified: modified(h),
		})

		for _, str := range h.SMap() {
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
	_, err := os.Stat(bag.path)

	old := !errors.Is(err, os.ErrNotExist)

	if bag.opts.Mode != types.NONE {
		var err error

		bag.file, err = os.OpenFile(bag.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)

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

	if len(bag.opts.Sign) > 0 {
		imp = hmac.New(sha256.New, []byte(bag.opts.Sign))
	} else {
		imp = sha256.New()
	}

	buf, err := os.ReadFile(bag.path)

	if err != nil {
		log.Println(err)
		return
	}

	imp.Write(buf)

	sum := base64.StdEncoding.EncodeToString(imp.Sum(nil))

	err = os.WriteFile(bag.path+".sig", []byte(sum), 0600)

	if err != nil {
		log.Println(err)
	}

	return
}

func modified(h *heap.Heap) time.Time {
	mt := time.Now().UTC()

	if h.Type == types.Regular {
		fi, err := os.Stat(h.Name)

		if err == nil {
			mt = fi.ModTime()
		} else {
			log.Println(err)
		}
	}

	return mt
}
