package widgets

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
)

const timeout = 1

type Overlay struct {
	base
	m   sync.RWMutex
	ch  chan message
	msg *message
}

type message struct {
	v string
	s tcell.Style
}

func NewOverlay(ctx *opt.Context) *Overlay {
	return &Overlay{
		base: base{ctx},

		ch: make(chan message, 64),
	}
}

func (o *Overlay) Render(x, y, w, _ int) {
	o.m.RLock()
	defer o.m.RUnlock()

	if o.msg != nil {
		o.print(x, y, fmt.Sprintf("%-*s", w, o.msg.v), o.msg.s)
	}
}

func (o *Overlay) SendError(err string) {
	o.ch <- message{err, themes.Overlay0}
}

func (o *Overlay) SendInfo(msg string) {
	o.ch <- message{msg, themes.Overlay1}
}

func (o *Overlay) Listen() {
	for {
		select {
		case msg := <-o.ch:
			o.m.Lock()
			o.msg = &msg
			o.m.Unlock()

		case <-time.After(timeout * time.Second):
			o.m.Lock()
			o.msg = nil
			o.m.Unlock()
		}

		o.ctx.ForceRender()
	}
}

func (o *Overlay) Close() {
	close(o.ch)
}
