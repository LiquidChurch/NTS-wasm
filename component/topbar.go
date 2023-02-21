package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func TopBar() *topBar {
	return &topBar{}
}

type topBar struct {
	app.Compo

	rightSide []app.UI
}

func (t *topBar) Right(v []app.UI) *topBar {
	t.rightSide = v
	return t
}

func (t *topBar) Render() app.UI {
	return app.Header().
		Body(
			app.Nav().
				Class("bg-topbar flex flex-row flex-wrap justify-between").
				Body(
					app.Div().
						Class("inline-flex h-12").
						Body(
							app.Img().
								Src("/web/asset/Liquid.svg").
								Class("my-auto pl-4 h-10"),
							app.H1().
								Class("my-auto px-4 text-2xl font-bold text-primary").
								Text("Night to Shine"),
						),
					app.Div().
						Class("flex-auto text-center place-self-center").
						Body(
							app.Range(t.rightSide).Slice(func(j int) app.UI {
								return t.rightSide[j]
							}),
						),
				),
		)
}
