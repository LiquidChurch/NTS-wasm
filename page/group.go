package page

import (
	"sort"
	"strings"
	"strconv"
	"nts/rockapi"
	"nts/component"
	"nts/element"
	"nts/modal"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Group struct {
	app.Compo
	groupId			int
	group 			*rockapi.GroupDetail
	groupMembers 	[]rockapi.GroupMemberDetail
	occurrence		*rockapi.OccurrenceDetail
	attendances 	[]*rockapi.AttendanceDetail
}

func (g *Group) OnMount(ctx app.Context) {
	g.updateGroup(ctx)
}


func (g *Group) OnAppUpdate(ctx app.Context) {
	updateAvailable := ctx.AppUpdateAvailable() 

	if updateAvailable == true {
		ctx.Reload()
	}	
}

func (g *Group) updateGroup(ctx app.Context) {
	path := strings.TrimPrefix(ctx.Page().URL().Path, "/v/")
	g.groupId, _ = strconv.Atoi(path)

	g.group = rockapi.LoadGroupById(ctx, g.groupId, 0)
	g.occurrence = rockapi.LoadOccurrenceByGroupId(ctx, g.groupId)

	if (g.group != nil) {
		ctx.NewActionWithValue("group", *g.group)
	}

	members := rockapi.LoadGroupMembersByGroupId(ctx, g.groupId)
	if (g.occurrence != nil) {
		g.attendances = rockapi.LoadAttendancesByOccurrenceId(ctx, g.occurrence.Id)
	}
	if (len(members) != 0) {
		for i, member := range members {
			personId := member.PersonId
			members[i].Person = *rockapi.LoadPersonById(ctx, personId)
			members[i].Attendance.DidAttend = false
			if g.occurrence != nil {
				for _, attendance := range g.attendances {
					if attendance.PersonAliasId == members[i].Person.PrimaryAliasId {
						attendanceId := attendance.Id
						if attendance.DidAttend == true {
							members[i].Attendance.DidAttend = true
							members[i].Attendance.Id = attendanceId
						}
					} 
				}
			}
			g.groupMembers = append(g.groupMembers, *members[i])
		}
		sort.Slice(g.groupMembers, func(i, j int) bool {
			return g.groupMembers[i].Person.LastName < g.groupMembers[j].Person.LastName
		  })
		ctx.NewActionWithValue("groupMembers", (g.groupMembers))
	}

}

func (g *Group) Render() app.UI {
	return app.Main().
	Class("").
	Body(
		component.TopBar(),
		app.Div().
			Class("p-4").
			Body(
				app.Div().
					Class("relative flex flex-col min-w-0 break-words border-gray-200 border-2 shadow-md rounded-2xl m-2 p-2 bg-clip-border").
					Body(
						element.Group(),
						element.GroupMembers(),
					),
				&modal.Handler{},
				),
	)
}
