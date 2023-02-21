package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Button() *button {
	return &button{}
}

type button struct {
	app.Compo
	label 	string
	width 	string
	color 	string
	href	string
	onClick app.EventHandler
}

func (b *button) Label(v string) *button {
	b.label = v
	return b
}

func (b *button) Width(v string) *button {
	b.width = v
	return b
}

func (b *button) Color(v string) *button {
	b.color = v
	return b
}

func (b *button) OnClick(v app.EventHandler) *button {
	b.onClick = v
	return b
}

func (b *button) Href(v string) *button {
	b.href = v
	return b
}

func (b *button) Render() app.UI {
	widthClass := " min-w-fit"
	bgColorClass := " bg-primary"

	switch b.width {
	case "small":
		widthClass = widthClass + " w-20"
	case "medium":
		widthClass = widthClass + " w-32"
	case "large":
		widthClass = widthClass + " w-44"
	default:
		widthClass = widthClass + " w-20"
	}

	switch b.color {
		case "primary":
			bgColorClass = " bg-blue-600"
		case "alert":
			bgColorClass = " bg-red-600"
		case "secondary":
			bgColorClass = " bg-green-600"
		case "grey":
			bgColorClass = " bg-grey-600"
		default:
			bgColorClass = " bg-blue-600"
		}

	return app.If(b.href == "", 
					app.Button().
						TabIndex('0').
						Class(widthClass + bgColorClass + " py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none mr-2").
						Text(b.label).
						OnClick(b.onClickHandler),
				).Else(
					app.A().
						TabIndex('0').
						Class(widthClass + bgColorClass + " py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none mr-2").
						Text(b.label).
						Href(b.href),
				)
}

func (b *button) onClickHandler(ctx app.Context, e app.Event) {
	if b.onClick != nil {
		b.onClick(ctx, e)
	}
}
