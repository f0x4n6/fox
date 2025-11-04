package widgets

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/v3/internal/opt"
	"github.com/cuhsat/fox/v3/internal/opt/ui/themes"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
)

const timeout = 3

type Overlay struct {
	base
	m    sync.RWMutex
	buf  *message
	msgs chan message
}

type message struct {
	v string
	s tcell.Style
}

func NewOverlay(state *opt.State) *Overlay {
	return &Overlay{
		base: base{state},
		msgs: make(chan message, 64),
	}
}

func (o *Overlay) Render(x, y, w, _ int) {
	o.m.RLock()
	defer o.m.RUnlock()

	if o.buf != nil {
		o.print(x, y, fmt.Sprintf("%-*s", w, o.buf.v), o.buf.s)
	}
}

func (o *Overlay) SendError(err string) {
	o.msgs <- message{err, themes.Overlay0}
}

func (o *Overlay) SendInfo(msg string) {
	o.msgs <- message{msg, themes.Overlay1}
}

func (o *Overlay) Listen() {
	for {
		select {
		case log := <-sys.Logs:
			o.SendError(log)

		case msg := <-o.msgs:
			o.m.Lock()
			o.buf = &msg
			o.m.Unlock()

		case <-time.After(timeout * time.Second):
			o.m.Lock()
			o.buf = nil
			o.m.Unlock()
		}

		o.state.ForceRender()
	}
}
