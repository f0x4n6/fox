package widgets

import (
	"math"

	"github.com/cuhsat/fox/internal/opt/ui/themes"
	"github.com/cuhsat/fox/internal/pkg/sys/fs"
)

func (v *View) listRender(p *panel) {
	w := fs.InfoWidth

	page := v.listPage()

	// print list
	for y, item := range page {
		v.print(p.X+0, p.Y+y, item.Info, themes.Subtext0)
		v.print(p.X+w, p.Y+y, item.Text, themes.Terminal)
	}

	// print select
	v.print(p.X+w, p.Y+(v.line), v.listItem().Text, themes.Subtext2)

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

func (v *View) listPage() []fs.Item {
	items := v.list.Load().([]fs.Item)

	return items[min(v.off, len(items)-1):]
}

func (v *View) listItem() fs.Item {
	page := v.listPage()

	return page[min(v.line, len(page)-1)]
}
