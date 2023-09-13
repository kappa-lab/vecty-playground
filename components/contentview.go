package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ContentView struct {
	vecty.Core
}

func NewContentView() *ContentView {
	return &ContentView{}
}
func (v *ContentView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("d-flex", "flex-column", "align-items-stretch", "bg-white"),
		),
		elem.Paragraph(vecty.Text("Home Contents")),
	)
}
