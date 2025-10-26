package widgets

import (
	"math"

	"github.com/cuhsat/fox/internal/opt/ui/adapter"
	"github.com/cuhsat/fox/internal/opt/ui/themes"
)

func (v *View) listRender(p *panel) {
	w := v.adapter.Width()

	page := v.listPage()

	// print list
	for y, node := range page {
		// print data
		v.print(p.X, p.Y+y, node.Data, themes.Subtext0)

		// print text
		v.print(p.X+w, p.Y+y, node.Text, themes.Terminal)
	}

	// print select
	v.print(p.X+w, p.Y+(v.line), page[v.line].Text, themes.Subtext2)

	// print scrollbar
	for y := range p.H {
		v.print(p.W-1, p.Y+y, "│", themes.Subtext1)
	}

	// render scrollbar
	if len(page) > 0 {
		scrollY := int(math.Round((float64(v.line+1) / float64(len(page))) * float64(p.H-1)))

		// fix zero position
		if v.line == 0 {
			scrollY = 0
		}

		// vertical scrollbar
		v.state.Root.SetContent(p.W-1, p.Y+scrollY, '┃', nil, themes.Terminal)
	}
}

func (v *View) listPage() []adapter.Node {
	nodes := v.list.Load().([]adapter.Node)

	return nodes[v.off:]
}
