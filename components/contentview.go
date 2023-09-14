package components

import (
	"context"
	"net/url"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/kappa-lab/vecty-playground/dispatcher"
)

type ContentView struct {
	vecty.Core
	url url.URL
}

func NewContentView(url url.URL) *ContentView {
	v := &ContentView{url: url}
	dispatcher.Register(v.OnUpdateURL)
	return v
}
func (v *ContentView) OnUpdateURL(url url.URL) {
	v.url = url
	vecty.Rerender(v)
}

func (v *ContentView) Render() vecty.ComponentOrHTML {
	var content vecty.ComponentOrHTML = elem.Div()

	title := v.url.Fragment
	switch v.url.Fragment {
	case "home":
		content = v.home()
	case "products":
		content = v.products()
	case "orders":
		content = v.orders()
	case "customers":
		content = v.customers(context.Background())
	case "":
		title = "MY SHOP"
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
			vecty.Text(title),
		),
		content,
	)
}

func (v *ContentView) home() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Paragraph(vecty.Text("SHOP-ID: 1")),
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

func (v *ContentView) customers(ctx context.Context) vecty.ComponentOrHTML {
	return NewCustomersView().Load(ctx)
}
