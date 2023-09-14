package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ContentView struct {
	vecty.Core
	Title string
}

func NewContentView() *ContentView {
	return &ContentView{Title: "Home"}
}
func (v *ContentView) Render() vecty.ComponentOrHTML {
	var content vecty.ComponentOrHTML = elem.Div()

	switch v.Title {
	case "products":
		content = v.products()
	case "orders":
		content = v.orders()
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("d-flex", "flex-column", "align-items-stretch", "bg-white"),
			vecty.Style("margin", "80px"),
		),
		elem.Heading1(
			vecty.Markup(
				vecty.Class("text-uppercase"),
			),
			vecty.Text(v.Title),
		),
		content,
	)
}

func (v *ContentView) products() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Table(
			vecty.Markup(
				vecty.Class("table", "table-sm", "table-bordered"),
			),
			elem.TableHead(
				elem.TableRow(
					elem.TableHeader(vecty.Text("PRODUCT-ID")),
					elem.TableHeader(vecty.Text("NAME")),
					elem.TableHeader(vecty.Text("PRICE")),
				),
			),
			elem.TableBody(
				elem.TableRow(
					elem.TableData(vecty.Text("1")),
					elem.TableData(vecty.Text("bag")),
					elem.TableData(vecty.Text("$100")),
				),
			),
			elem.TableBody(
				elem.TableRow(
					elem.TableData(vecty.Text("2")),
					elem.TableData(vecty.Text("pen")),
					elem.TableData(vecty.Text("$110")),
				),
			),
		),
	)
}
func (v *ContentView) orders() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Table(
			vecty.Markup(
				vecty.Class("table", "table-sm", "table-bordered"),
			),
			elem.TableHead(
				elem.TableRow(
					elem.TableHeader(vecty.Text("ORDER-ID")),
					elem.TableHeader(vecty.Text("PRODUCT-ID")),
					elem.TableHeader(vecty.Text("DATE")),
				),
			),
			elem.TableBody(
				elem.TableRow(
					elem.TableData(vecty.Text("1001")),
					elem.TableData(vecty.Text("1")),
					elem.TableData(vecty.Text("2023-01-01 10:00:00")),
				),
			),
			elem.TableBody(
				elem.TableRow(
					elem.TableData(vecty.Text("1002")),
					elem.TableData(vecty.Text("2")),
					elem.TableData(vecty.Text("2023-01-01 11:01:01")),
				),
			),
		),
	)
}
