package widgets

import (
	"fmt"
	"strings"
	"sync/atomic"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/opt"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type Position int

const (
	Before Position = iota
	After
)

const Cursor = tcell.CursorStyleBlinkingBar

type Prompt struct {
	base
	lock      atomic.Bool
	value     atomic.Value
	offset    atomic.Int32
	cursor    atomic.Int32
	cursorEnd atomic.Int32
	cursorMax atomic.Int32
}

func NewPrompt(state *opt.State) *Prompt {
	p := Prompt{base: base{state}}

	p.lock.Store(true)
	p.value.Store("")
	p.offset.Store(0)
	p.cursor.Store(0)
	p.cursorEnd.Store(0)
	p.cursorMax.Store(0)

	return &p
}

func (p *Prompt) Render(hs *heapset.HeapSet, x, y, w, _ int) int {
	var hl, hc int
	var fs []string

	if hs != nil {
		_, heap := hs.Heap()
		hl = heap.LastFilter().Len()
		hc = heap.LastFilter().Context.Size()
		fs = heap.Patterns()
	}

	lb := len(fs)
	fs = p.fitFilters(fs)
	la := len(fs)

	m := p.fmtMode()
	f := p.fmtFilters(fs, la < lb)
	i := p.fmtInput(fs)

	s := p.fmtStatus(hl, hc)

	// render blank line
	p.blank(x, y, w, themes.Surface0)

	// render mode
	p.print(x, y, m, themes.Surface3)

	// skip the rest in static modes
	if p.state.Mode().Static() {
		return 1
	}

	ml := text.Len(m)
	fl := text.Len(f)
	sl := text.Len(s)

	x += ml

	// render filter
	if p.state.Mode().Filter() {
		p.print(x, y, f, themes.Surface1)
	}

	x += fl

	// render input
	if len(i) > 1 {
		p.print(x, y, i, themes.Surface1)
	}

	// render status
	p.print(w-sl, y, s, themes.Surface1)

	// calculate cursor position
	vl := text.Len(p.value.Load().(string))
	cm := max(w-(ml+fl+sl), 0)
	c := int(p.cursor.Load())
	o := int(p.offset.Load())
	d := 0

	// set space for first input
	if fl == 0 {
		d = 1
	}

	p.cursorEnd.Store(int32(vl - o))
	p.cursorMax.Store(int32(cm - 1))

	if p.state.Mode().Prompt() && !p.Locked() {
		p.state.Root.ShowCursor(x+d+c, y)
	} else {
		p.state.Root.HideCursor()
	}

	return 1
}

func (p *Prompt) MoveStart() {
	p.cursor.Store(0)
	p.offset.Store(0)
}

func (p *Prompt) MoveEnd() {
	ce := p.cursorEnd.Load()
	cm := p.cursorMax.Load()

	vl := int32(text.Len(p.value.Load().(string)))

	p.cursor.Store(min(ce, cm))
	p.offset.Store(max(0, vl-cm))
}

func (p *Prompt) Move(d int) {
	c := p.cursor.Add(int32(d))
	o := p.offset.Load()
	ce := p.cursorEnd.Load()
	cm := p.cursorMax.Load()

	p.cursor.Store(min(max(c, 0), ce, cm))

	vl := int32(text.Len(p.value.Load().(string)))

	// scroll past start
	if c < 0 && o > 0 {
		p.offset.Add(-1)
	}

	// scroll past end
	if c > cm && o+c <= vl {
		p.offset.Add(1)
	}
}

func (p *Prompt) Lock(b bool) {
	p.lock.Store(b)
}

func (p *Prompt) Locked() bool {
	return p.lock.Load()
}

func (p *Prompt) AddRune(r rune) {
	if !utf8.ValidRune(r) {
		return
	}

	v := p.value.Load().(string)
	o := int(p.offset.Load())
	c := int(p.cursor.Load())
	cm := int(p.cursorMax.Load())

	if p.Locked() {
		return
	}

	p.value.Store(v[:o+c] + string(r) + v[o+c:])

	// move cursor if there is space left
	if c < cm {
		p.cursorEnd.Add(1)
		p.Move(1)
	} else {
		p.offset.Add(1)
	}
}

func (p *Prompt) DelRune(po Position) {
	v := p.value.Load().(string)
	o := int(p.offset.Load())
	c := int(p.cursor.Load())

	if p.Locked() || len(v) <= 0 {
		return
	}

	lv := text.Len(v)

	p.cursorEnd.Add(-1)

	switch po {
	case Before:
		p.value.Store(v[:max(o+c-1, 0)] + v[o+c:])

		if o > 0 {
			p.offset.Add(-1)
		} else {
			p.Move(-1)
		}
	case After:
		p.value.Store(v[:o+c] + v[min(o+c+1, lv):])
	}
}

func (p *Prompt) ReadLine() (s string) {
	if p.Locked() || p.cursorMax.Load() == 0 {
		return
	}

	s = p.GetValue()

	p.SetValue("")

	return
}

func (p *Prompt) GetValue() string {
	return p.value.Load().(string)
}

func (p *Prompt) SetValue(s string) {
	if p.Locked() || !utf8.ValidString(s) {
		return
	}

	cm := int(p.cursorMax.Load())
	o := max(len(s)-cm, 0)
	c := min(len(s)-o, len(s))

	p.value.Store(s)

	p.cursor.Store(int32(c))
	p.offset.Store(int32(o))
}

func (p *Prompt) fmtMode() string {
	return fmt.Sprintf(" %s ", p.state.Mode())
}

func (p *Prompt) fitFilters(fs []string) []string {
	w, _ := p.state.Root.Size()

	for {
		l := text.Len(p.fmtFilters(fs, false))

		if len(fs) == 1 || l <= int(float32(w)/1.5) {
			return fs
		}

		fs = fs[1:]
	}
}

func (p *Prompt) fmtFilters(fs []string, fit bool) string {
	var sb strings.Builder

	if p.state.Mode().Filter() {
		for i, f := range fs {
			if fit && i == 0 {
				sb.WriteRune(' ')
				sb.WriteRune(p.state.Icon.Abr)
			}
			sb.WriteRune(' ')
			sb.WriteString(f)
			sb.WriteRune(' ')
			sb.WriteRune(p.state.Icon.Grep)
		}

		// add space after filters
		if len(fs) > 0 {
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}

func (p *Prompt) fmtInput(fs []string) string {
	var sb strings.Builder

	if v, ok := p.value.Load().(string); ok {
		// add space before input in all modes
		if (!p.state.Mode().Filter() || len(fs) == 0) && len(v) > 0 {
			sb.WriteRune(' ')
		}

		sb.WriteString(text.Trim(
			v,
			int(p.offset.Load()),
			int(p.cursorMax.Load()),
		))

		// add space after input
		sb.WriteRune(' ')
	}

	return sb.String()
}

func (p *Prompt) fmtStatus(l, c int) string {
	var sb strings.Builder

	if c > 0 {
		sb.WriteString(fmt.Sprintf(" %d%c%d ", l, p.state.Icon.Size, c))
	} else {
		sb.WriteString(fmt.Sprintf(" %d ", l))
	}

	if p.state.IsNavi() {
		sb.WriteRune('N')
	} else {
		sb.WriteRune(p.state.Icon.None)
	}

	if p.state.IsWrap() {
		sb.WriteRune('W')
	} else {
		sb.WriteRune(p.state.Icon.None)
	}

	if p.state.IsFollow() {
		sb.WriteRune('T')
	} else {
		sb.WriteRune(p.state.Icon.None)
	}

	sb.WriteRune(' ')

	return sb.String()
}
