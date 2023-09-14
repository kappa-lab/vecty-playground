package main

import (
	"fmt"
	"net/url"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/kappa-lab/vecty-playground/components"
)

func main() {
	href := js.Global().Get("location").Get("href")
	url, err := url.Parse(fmt.Sprintf("%s", href))

	if err != nil {
		panic(err)
	}

	vecty.SetTitle("My Shop")

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css")
	vecty.AddStylesheet("https://getbootstrap.jp/docs/5.0/examples/sidebars/sidebars.css")

	vecty.RenderBody(components.NewPageView(*url))

}
