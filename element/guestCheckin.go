package element

import (
	"log"
	"strconv"
	"nts/component"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func GuestCheckin() *guestCheckin{
	return &guestCheckin{}
}

type guestCheckin struct {
	app.Compo

	groupId 		int
	attendance 		*rockapi.AttendanceDetail
	guestId			rockapi.GuestId
	attended		bool
	editing			bool
	infoSet		bool
}

func (g *guestCheckin) OnMount(ctx app.Context) {
	g.infoSet = false
	g.attended = false
	g.editing = false
	g.groupId = 263313
	ctx.Handle("guestCheckin", g.setInfo)
	ctx.Handle("editRedirect", g.editLink)
	g.infoSet = true
}

func (g *guestCheckin) setInfo(ctx app.Context, a app.Action) {
	g.guestId = interface{}(a.Value).(rockapi.GuestId)
	occurrance := rockapi.LoadOccurrenceByGroupId(ctx, g.groupId)
	if occurrance != nil {
		attendance := rockapi.LoadAttendanceByPersonAliasId(ctx, g.guestId.PersonAliasId, occurrance.Id)
		if attendance != nil {
			g.attended = attendance.DidAttend
		} 
	}
}

func (g *guestCheckin) editLink(ctx app.Context, a app.Action) {
	log.Println("Edited Guest Id", g.guestId.GuestId)
	ctx.Navigate("https://my.liquidchurch.com/page/852?RegistrantId=" + strconv.Itoa(g.guestId.GuestId))
}

func (g *guestCheckin) addAttendance(ctx app.Context, e app.Event) {
	attendance := rockapi.AddAttendanceByGroupId(ctx, g.guestId.PersonId, g.groupId)
	g.attended = attendance.DidAttend
}

func (g *guestCheckin) openEdit(ctx app.Context, e app.Event) {
	log.Println("Guest Info Edit Triggered")
	g.editing = true
	ctx.NewAction("editRedirect")
}

func (g *guestCheckin) reload(ctx app.Context, e app.Event) {
	log.Println("Guest Info Refresh Triggered")
	ctx.Reload()
}

func (g *guestCheckin) goHome(ctx app.Context, e app.Event) {
	ctx.NewActionWithValue("readAppKey", "")
	ctx.Navigate("/")
}

func (g *guestCheckin) Render() app.UI {
	var	infoSet bool
	if g.infoSet {
		log.Println("Guest Checkin Loaded")
		infoSet = true
	} else {
		log.Println("Guest Checkin Not Loaded")
		return app.Div().Body()
	}

	return app.If(infoSet == true,
			app.If(g.attended == true,
				app.Div().Body(
					component.Button().
						Width("medium").
						Color("primary").
						Label("Back").
						OnClick(g.goHome),
					app.Button().
						TabIndex('0').
						Class("w-44 bg-gray-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue active:shadow-none mr-2").
						Text("Checked In"),
				),
			).ElseIf(g.editing == true,
				app.Div().Body(
					component.Button().
						Width("medium").
						Color("primary").
						Label("Back").
						OnClick(g.goHome),
					app.Button().
						TabIndex('0').
						Class("w-44 bg-green-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue active:shadow-none mr-2").
						Text("Refresh").
						OnClick(g.reload),
				),
			).Else(
				app.Div().Body(
					component.Button().
						Width("medium").
						Color("primary").
						Label("Back").
						OnClick(g.goHome),
					component.Button().
						Width("large").
						Color("secondary").
						Label("Check-In").
						OnClick(g.addAttendance),
					component.Button().
						Width("medium").
						Color("alert").
						Label("Edit").
						OnClick(g.openEdit),
					),
				),
	)
}
