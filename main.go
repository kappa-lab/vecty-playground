package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	vecty.SetTitle("Mypage")
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello World"),
	)
}
