package element

import (
	"log"
	"time"
	"nts/rockapi"
	"nts/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func GroupMember() *groupMember {
	return &groupMember{}
}

type groupMember struct {
	app.Compo
	groups []rockapi.GroupMemberDetail
}

func (g *groupMember) OnMount(ctx app.Context) {
	ctx.Handle("groupMember", g.setGroup)
}

func (g *groupMember) setGroup(ctx app.Context, a app.Action) {
	groups := interface{}(a.Value).([]*rockapi.GroupMemberDetail)

	for _, value := range groups {
		if value.Group.Attributes.CheckinDate.ValueFormatted == "" {
			value.Group.Attributes.CheckinDate.ValueFormatted = "2/10/2023"
		}
		g.groups = append(g.groups, *value)
	}
}

func (g *groupMember) Render() app.UI {

	log.Println(g.groups)


	todayDate := time.Now().Format("1/2/2006")
	log.Println(todayDate)
	var groupsSet bool
	if len(g.groups) == 0 {
		groupsSet = false
	} else {
		groupsSet = true
	}
	
	return app.Div().
		Class("").
		Body(
			app.Div().
				Body(
					app.If(groupsSet == true,
						component.Header().Level(3).Text("Thank you for serving at Night to Shine. You are on the following Team."),
						app.Range(g.groups).
							Slice(func(i int) app.UI {
								group := g.groups[i]
								flag := ""
								if (group.GroupRoleId == 119) { flag = "Leader"}
								return app.If(group.Group.Attributes.CheckinDate.ValueFormatted == todayDate,
									app.Div().
									Class("p-2 mb-4").
									Body(
										component.Card().Label(group.Group.Name).Flag(flag).
											Content(
												[]app.UI{
													app.Div().
														Class("px-4 pb-2").
														Body(
															component.Header().Level(4).Text("Here is the What to Expect document outlining arrival details and other important information for the night."),
															app.A().
																Class("text-xl font-normal leading-normal text-blue-800 pl-4").
																Text("What to Expect").
																Href(group.Group.Attributes.WhattoExpect.Value),
															component.Header().Level(4).Text("Please refer to the Map, marked with a Gold Crown, for your team's Huddle location."),
																app.A().
																	Class("text-xl font-normal leading-normal text-blue-800 pl-4").
																	Text("Huddle Location Map").
																	Href(group.Group.Attributes.Map.Value),
														),
														VolCheckin().Group(group),
												}),
									),
								)
							}),
					),
				),
			)
}