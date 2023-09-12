package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

type PageView struct {
	vecty.Core
	buttonLabel string
	message     string
	messages    []string
}

func NewPageView() *PageView {
	return &PageView{buttonLabel: "Clear"}
}

func (p *PageView) reset() *PageView {
	p.messages = []string{}
	p.message = ""
	return p
}
func (p *PageView) Render() vecty.ComponentOrHTML {

	messages := []vecty.MarkupOrChild{}
	for _, v := range p.messages {
		messages = append(messages, elem.ListItem(vecty.Text(v)))
	}

	return elem.Body(
		vecty.Markup(
			style.Margin(style.Px(10)),
		),
		elem.Script(
			vecty.Text("console.log('Ready...');"),
			vecty.Text("function log(val) {console.log(val)};"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("display-2"),
			),
			elem.Paragraph(
				vecty.Text("HelloWorld"),
			),
		),

		elem.Div(
			elem.Form(
				vecty.Markup(
					event.Submit(func(e *vecty.Event) {
						p.messages = append(p.messages, p.message)
						p.message = ""
						vecty.Rerender(p)
					}).PreventDefault(),
				),

				elem.Input(
					vecty.Markup(
						prop.Placeholder("Enter Message"),
						prop.Autofocus(true),
						prop.Value(p.message),
						event.Input(func(e *vecty.Event) {
							p.message = e.Target.Get("value").String()
						}),
					),
				),
			),
		),

		elem.Div(
			messages...,
		),

		elem.Div(
			elem.Form(
				vecty.Markup(
					event.Click(func(e *vecty.Event) {
						vecty.Rerender(p.reset())
					}),
				),

				elem.Button(
					vecty.Text(p.buttonLabel), // initial textarea text.
				),
			),
		),
	)
}
