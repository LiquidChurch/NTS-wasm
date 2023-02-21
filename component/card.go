package component

import (
	"log"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Card() *card {
	return &card{}
}

type card struct {
	app.Compo
	label 		string
	flag 		string

	content 	[]app.UI
	onClick 	app.EventHandler
	member		rockapi.GroupMemberDetail
}

func (c *card) Label(v string) *card {
	c.label = v
	return c
}

func (c *card) Flag(v string) *card {
	c.flag = v
	return c
}

func (c *card) Content(v []app.UI) *card {
	c.content = v
	return c
}

func (c *card) Member(v rockapi.GroupMemberDetail) *card {
	c.member = v
	return c
}

func (c *card) OnClick(v app.EventHandler) *card {
	c.onClick = v
	return c
}

func (c *card) onClickHandler(ctx app.Context, e app.Event) {
	log.Println("clicked", c.label)
	ctx.SetState("checkinModal", c.member)

	if c.onClick != nil {
		c.onClick(ctx, e)
	}
}

func (c *card) Render() app.UI {
	return app.Div().
			   Class("relative flex flex-col min-w-0 break-words border-gray-200 border-2 shadow-md rounded-2xl bg-clip-border").
			   Body(
				   	app.Div().
					   	Class("flex flex-row").
						Body(
							app.H4().
								Class("flex flex-auto text-xl font-normal leading-normal p-4 text-blue-800").
								Text(c.label),
							app.H5().
								Class("flex text-base font-normal text-primary p-4 content-center").
								Text(c.flag),
					   	),
					app.Range(c.content).Slice(func(j int) app.UI {
						return c.content[j]
					}),
			   ).OnClick(c.onClickHandler)
}