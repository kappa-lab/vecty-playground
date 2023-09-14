package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type CustomersJSON struct {
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type CustomersView struct {
	vecty.Core
	Data CustomersJSON
}

func NewCustomersView() *CustomersView {
	v := &CustomersView{}
	return v
}

func (v *CustomersView) Load(ctx context.Context) vecty.ComponentOrHTML {
	go func() { // HTTPリクエストを送信する場合は、goroutine化する必要がある
		endpoint := "/api/customers.json"
		var (
			resp *http.Response
			err  error
		)
		resp, err = http.Get(endpoint)
		if err != nil {
		}
		var data CustomersJSON
		err = json.NewDecoder(resp.Body).Decode(&data)

		vecty.RenderInto("#customers-inner", &CustomersView{Data: data})
	}()

	return elem.Div(
		vecty.Markup(
			vecty.Attribute("id", "customers-inner"),
		),
		elem.Paragraph(vecty.Text("Loading...")),
	)
}

func (v *CustomersView) Render() vecty.ComponentOrHTML {
	tableInner := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("table", "table-sm", "table-bordered"),
		),
		elem.TableHead(
			elem.TableRow(
				elem.TableHeader(vecty.Text("USER-ID")),
				elem.TableHeader(vecty.Text("NAME")),
			),
		),
	}

	for _, v := range v.Data.Data {
		tableInner = append(tableInner,
			elem.TableBody(
				elem.TableRow(
					elem.TableData(vecty.Text(fmt.Sprint(v.ID))),
					elem.TableData(vecty.Text(v.Name)),
				),
			),
		)
	}

	return elem.Div(
		vecty.Markup(
			vecty.Attribute("id", "customers-inner"),
		),
		elem.Table(tableInner...))
}
