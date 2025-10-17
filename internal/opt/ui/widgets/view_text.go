package widgets

import (
	"math"
	"strings"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/page"
)

func (v *View) textRender(p *panel) {
	pg := page.Text(&page.Context{
		Heap:    v.heap,
		Context: v.heap.HasContext(),
		Pinned:  v.state.IsPinned(),
		Navi:    v.state.IsNavi(),
		Wrap:    v.state.IsWrap(),
		Space:   v.state.Space(),
		Nr:      v.nr,
		X:       v.delta.X,
		Y:       v.delta.Y,
		W:       p.W,
		H:       p.H,
	})

	v.fmap = pg.FMap

	// set page bounds
	v.last.X = max(pg.W-1, 0)
	v.last.Y = max(pg.H-1, 0)

	// set preserved line
	if v.nr > 0 {
		y, _ := v.fmap.Find(v.nr)
		v.delta.Y = min(y, v.last.Y)
	}

	// reset
	v.nr = 0

	i := 0

	chat := v.heap.Type == types.Chat
	ps1 := string(v.state.Icon.Ps1)

	// render lines
	var color tcell.Style

	for line := range pg.Lines {
		lineX := p.X
		lineY := p.Y + i

		i++

		// context separators
		if line.Nr == "--" {
			v.print(lineX, lineY, strings.Repeat("―", p.W), themes.Subtext1)
			continue
		}

		// line number
		if v.state.IsNavi() {
			v.print(lineX, lineY, line.Nr, themes.Subtext0)
			lineX += len(line.Nr) + 1
		}

		// text input
		if len(line.Str) > 0 {
			if chat && strings.HasPrefix(line.Str, ps1) {
				color = themes.Subtext2
			} else if i == 1 && v.state.IsPinned() {
				color = themes.Subtext0
			} else {
				color = themes.Terminal
			}

			v.print(lineX, lineY, line.Str, color)
		}
	}

	// render parts on top
	for part := range pg.Parts {
		partX := p.X + part.X
		partY := p.Y + part.Y

		if v.state.IsNavi() {
			partX += pg.N + 1
		}

		// part input
		v.print(partX, partY, part.Str, themes.Subtext2)
	}

	// render scrollbars
	if v.state.IsNavi() {
		w := p.W - 1
		h := p.H - 1
		x := p.X
		y := p.Y

		scrollX := int(math.Round((float64(v.delta.X+1) / float64(v.last.X+1)) * float64(w-2)))
		scrollY := int(math.Round((float64(v.delta.Y+1) / float64(v.last.Y+1)) * float64(h-1)))

		// fix zero positions
		if v.delta.X == 0 {
			scrollX = 0
		}

		if v.delta.Y == 0 {
			scrollY = 0
		}

		for i := range w {
			v.state.Root.SetContent(x+i, y+h, '─', nil, themes.Subtext1)
		}

		for i := range h {
			v.state.Root.SetContent(x+w, y+i, '│', nil, themes.Subtext1)
		}

		v.state.Root.SetContent(x+w, y+h, '┘', nil, themes.Subtext1)

		// horizontal scrollbar
		v.state.Root.SetContent(x+scrollX+0, y+h, '━', nil, themes.Terminal)
		v.state.Root.SetContent(x+scrollX+1, y+h, '━', nil, themes.Terminal)

		// vertical scrollbar
		v.state.Root.SetContent(x+w, y+scrollY, '┃', nil, themes.Terminal)
	}
}

func (v *View) textGoto(s string) {
	var nr int

	switch s[0] {
	case '+':
		nr = (*v.fmap)[v.delta.Y].Nr + text.Int(s[1:])

	case '-':
		nr = (*v.fmap)[v.delta.Y].Nr - text.Int(s[1:])

	default:
		nr = text.Int(s)
	}

	if y, ok := v.fmap.Find(nr); ok {
		v.ScrollTo(v.delta.X, y)
	}
}
