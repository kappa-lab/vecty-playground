package main

import (
	"fmt"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/kappa-lab/vecty-playground/components"
)

func main() {
	location := js.Global().Get("location").Get("href")
	vecty.SetTitle("Mypage")

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css")

	vecty.RenderBody(components.NewPageView(fmt.Sprintf("%s", location)))

}
