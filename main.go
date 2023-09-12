package main

import (
	"github.com/hexops/vecty"
	"github.com/kappa-lab/vecty-playground/components"
)

func main() {

	vecty.SetTitle("Mypage")

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css")

	vecty.RenderBody(components.NewPageView())

}
