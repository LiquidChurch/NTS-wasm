package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ButtonGroup() *buttonGroup {
	return &buttonGroup{}
}

type buttonGroup struct {
	app.Compo

	buttonList []app.UI
}

func (b *buttonGroup) ButtonList(v []app.UI) *buttonGroup {
	b.buttonList = v
	return b
}

func (b *buttonGroup) Render() app.UI {

	return app.Div().
		Class("flex w-full items-center justify-center").
		Body(
			app.Range(b.buttonList).Slice(func(i int) app.UI {
				return b.buttonList[i]
			}),
		)
}