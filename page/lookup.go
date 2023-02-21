package page

import (
	"log"
	"strings"
	"nts/rockapi"
	"nts/component"
	"nts/element"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Lookup struct {
	app.Compo
	appKey string
	registrant *rockapi.RegistrantDetail
	groupMembers []*rockapi.GroupMemberDetail
	attendance *rockapi.AttendanceDetail
}

func (h *Lookup) OnMount(ctx app.Context) {
	h.updateGuid(ctx)
}

func (h *Lookup) OnAppUpdate(ctx app.Context) {
	updateAvailable := ctx.AppUpdateAvailable() 

	if updateAvailable == true {
		ctx.Reload()
	}	
}


func (h *Lookup) updateGuid(ctx app.Context) {
	h.appKey = strings.TrimPrefix(ctx.Page().URL().Path, "/")

	log.Println("Look up registrant")
	h.registrant = rockapi.LoadRegistrantByAppKey(ctx, h.appKey)

	if (h.registrant != nil) {
		log.Println("registrant found")
		ctx.NewActionWithValue("person", h.registrant.Person)
	}

	log.Println("Look up Groups", h.registrant.Person.Id)
	h.groupMembers = rockapi.LoadGroupMembersByPersonId(ctx, h.registrant.Person.Id, h.registrant.PersonAliasId)
	if (h.groupMembers != nil) {
		log.Println("groups found")
		ctx.NewActionWithValue("groupMember", h.groupMembers)
	}


}

func (h *Lookup) Render() app.UI {
	return app.Main().
	Class("").
	Body(
		component.TopBar(),
		app.Div().
			Class("p-4").
			Body(
				app.Div().
					Class("relative flex flex-col min-w-0 break-words border-gray-200 border-2 shadow-xl rounded-2xl m-2 p-2 bg-clip-border").
					Body(
						element.Person(),
						element.GroupMember(),
					),
				),
		component.NewGeoLocation(),
	)
}