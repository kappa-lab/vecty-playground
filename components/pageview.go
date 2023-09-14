package components

import (
	"net/url"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type PageView struct {
	vecty.Core
	url         url.URL
	sidebarView *SidebarView
	contentView *ContentView
}

func NewPageView(url url.URL) *PageView {
	contentView := NewContentView()
	contentView.Title = url.Fragment
	contentView.Title = url.Fragment
	return &PageView{
		url:         url,
		sidebarView: NewSidebarView(url),
		contentView: contentView,
	}
}

func (v *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			v.sidebarView,
			v.contentView,
			// elem.Div(
			// 	elem.Form(
			// 		vecty.Markup(
			// 			event.Click(func(e *vecty.Event) {
			// 				v.contentView.Content = "Click"
			// 				vecty.Rerender(v)
			// 			}),
			// 		),

			// 		elem.Button(
			// 			vecty.Text("BTN"), // initial textarea text.
			// 		),
			// 	),
			// ),
		),
	)
}
