package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Header() *header {
	return &header{}
}

type header struct {
	app.Compo
	level 	int
	text 	string
}

func (h *header) Level(v int) *header {
	h.level = v
	return h
}

func (h *header) Text(v string) *header {
	h.text = v
	return h
}

func (h *header) Render() app.UI {
	return app.Div().
			Body(app.If(h.level == 1,
				app.H1().
					Class("text-3xl font-normal leading-normal my-1 text-blue-800").
					Text(h.text),
			).ElseIf(h.level == 2,
				app.H2().
					Class("text-2xl font-normal leading-normal my-1 text-blue-800").
					Text(h.text),
			).ElseIf(h.level == 3,
				app.H3().
					Class("text-xl font-normal leading-normal my-1 text-primary").
					Text(h.text),
			).ElseIf(h.level == 4,
				app.H4().
					Class("text-base font-normal leading-normal my-1 text-primary").
					Text(h.text),
			),
			)
}