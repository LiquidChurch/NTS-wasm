package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func InfoCard() *infoCard {
	return &infoCard{}
}

type infoCard struct {
	app.Compo
	
	label	string
	content []app.UI
}

func (i *infoCard) Label(v string) *infoCard {
	i.label = v
	return i
}

func (i *infoCard) Content(v []app.UI) *infoCard {
	i.content = v
	return i
}

func (i *infoCard) Render() app.UI {

	return app.Div().
			Class("relative flex flex-col min-w-0 break-words border-gray-200 border-2 shadow-md rounded-2xl mx-2 my-4 p-2 bg-clip-border").
			Body(
				app.Label().Class().Text(i.label),
				app.Range(i.content).Slice(func(j int) app.UI {
					return i.content[j]
				}),
			)
}