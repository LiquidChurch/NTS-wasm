package element

import (
	"log"
	"nts/rockapi"
	"nts/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func GroupMembers() *groupMembers {
	return &groupMembers{}
}

type groupMembers struct {
	app.Compo
	
	groupMembers 	[]rockapi.GroupMemberDetail
	isLoaded		bool
}

func (g *groupMembers) OnMount(ctx app.Context) {
	g.isLoaded = false
	ctx.Handle("groupMembers", g.setMembers)
}

func (g *groupMembers) setMembers(ctx app.Context, a app.Action) {
	g.groupMembers = interface{}(a.Value).([]rockapi.GroupMemberDetail)
	g.isLoaded = true
}

func (g *groupMembers) openCheckinModal(ctx app.Context, e app.Event) {
	log.Println("open checkin modal")
	ctx.NewActionWithValue("modal", "checkinModal")
}


func (g *groupMembers) Render() app.UI {
	var groupSet bool
	if g.isLoaded == false {
		groupSet = false
	} else {
		groupSet = true
	}

	return app.Div().
	Class("").
	Body(
		app.If(groupSet == true,
			component.Header().Level(3).Text("Members"),
			app.Range(g.groupMembers).
				Slice(func(i int) app.UI {
					member := g.groupMembers[i]
					flag := "Not Checked In"
					if member.Attendance.DidAttend == true {
						flag = "Checked In"
					} 
					return component.Card().
								Label(member.Person.FullName).
								Flag(flag).
								Member(member).
								OnClick(g.openCheckinModal)
				}),
		),
	)
}