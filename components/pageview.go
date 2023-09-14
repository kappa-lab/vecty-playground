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
	return &PageView{
		url:         url,
		sidebarView: NewSidebarView(url),
		contentView: NewContentView(url),
	}
}

func (v *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			v.sidebarView,
			v.contentView,
		),
	)
}
