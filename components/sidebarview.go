package components

import (
	"net/url"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/kappa-lab/vecty-playground/dispatcher"
)

type SidebarView struct {
	vecty.Core
	url url.URL
}

func NewSidebarView(url url.URL) *SidebarView {
	v := &SidebarView{url: url}
	dispatcher.Register(v.OnUpdateURL)
	return v
}

func (v *SidebarView) OnUpdateURL(url url.URL) {
	v.url = url
	vecty.Rerender(v)
}

func (v *SidebarView) addItem(text string, fragment string) *vecty.HTML {
	state := "text-white"
	if v.url.Fragment == fragment {
		state = "active"
	}
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item")),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link", state),
				vecty.Attribute("aria-current", "page"),
				vecty.Attribute("href", "/#"+fragment),
				event.Click(func(e *vecty.Event) {
					v.url.Fragment = fragment
					dispatcher.Dispatch(v.url)
				}),
			),
			vecty.Text(text),
		),
	)
}
func (v *SidebarView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("d-flex", "flex-column", "p-3", "text-white", "bg-dark"),
			vecty.Style("width", "280px"),
			vecty.Style("height", "100vh"),
		),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("d-flex", "align-items-center", "mb-3", "mb-md-0", "me-md-auto", "text-white", "text-decoration-none"),
				vecty.Attribute("href", "/"),
			),
			elem.Span(
				vecty.Markup(vecty.Class("fs-4")),
				vecty.Text("My Shop"),
			),
		),
		elem.HorizontalRule(),
		elem.UnorderedList(
			vecty.Markup(vecty.Class("nav", "nav-pills", "flex-column", "mb-auto")),
			v.addItem("Home", "home"),
			v.addItem("Dashboard", "dashboard"),
			v.addItem("Orders", "orders"),
			v.addItem("Products", "products"),
			v.addItem("Customers", "customers"),
		),
		elem.HorizontalRule(),
		elem.Div(
			vecty.Markup(
				vecty.Class("text-center", "small"),
			),
			elem.Paragraph(vecty.Text("ver 0.0.1")),
		),
	)
}
