package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			vecty.Class("display-2"),
		),
		elem.Paragraph(
			vecty.Text("HelloWorld"),
		),
	)
}
