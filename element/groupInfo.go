package element

import (
	"log"
	"nts/rockapi"
	"nts/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Group() *group {
	return &group{}
}

type group struct {
	app.Compo
	group 		rockapi.GroupDetail
	isLoaded 	bool
}

func (g *group) OnMount(ctx app.Context) {
	g.isLoaded = false
	ctx.Handle("group", g.setGroup)
}

func (g *group) setGroup(ctx app.Context, a app.Action) {
	log.Println("set Group start")
	g.group = interface{}(a.Value).(rockapi.GroupDetail)
	g.isLoaded = true
}

func (g *group) Render() app.UI {
	log.Println("group element")
	var groupSet bool
	if g.isLoaded == false {
		groupSet = false
	} else {
		groupSet = true
	}

	log.Println("groupsSet", groupSet)

	return app.Div().
	Class("").
	Body(
		app.If(groupSet == true,
			component.Header().Level(3).Text("Group"),
			component.Header().Level(1).Text(g.group.Name)),
			component.Header().Level(3).Text("Group Documents:"),
			app.Div().
				Class("flex flex-col").
				Body(
				app.A().
					Class("text-xl font-normal leading-normal text-blue-800 pl-4").
					Text("What to Expect").
					Href(g.group.Attributes.WhattoExpect.Value),
				app.A().
					Class("text-xl font-normal leading-normal text-blue-800 pl-4").
					Text("Huddle Location Map").
					Href(g.group.Attributes.Map.Value),
				app.A().
					Class("text-xl font-normal leading-normal text-blue-800 pl-4").
					Text("Team Lead Information").
					Href("/teamlead"),	
				),
	)
}
