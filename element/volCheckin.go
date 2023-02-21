package element

import (
	"log"
	"strconv"
	"nts/component"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func VolCheckin() *volCheckin{
	return &volCheckin{}
}

type volCheckin struct {
	app.Compo
	
	group		rockapi.GroupMemberDetail
	locStatus	string
}

func (v *volCheckin) OnMount(ctx app.Context) {
	ctx.Handle("geoLocUpdate", v.geoLocStatus)
	ctx.GetState("geoLocation", &v.locStatus)
	log.Println("Location Status", v.locStatus)
}

func (v *volCheckin) geoLocStatus(ctx app.Context, a app.Action) {
	ctx.GetState("geoLocation", &v.locStatus)
	log.Println("Location Status", v.locStatus)
}

func (v *volCheckin) Group(w rockapi.GroupMemberDetail) *volCheckin {
	v.group = w
	return v
}

func (v *volCheckin) addAttendance(ctx app.Context, e app.Event) {
	attendance := rockapi.AddAttendanceByGroupId(ctx, v.group.PersonId, v.group.GroupId)
	v.group.Group.Attendance = *attendance
}

func (v *volCheckin) openGroup(ctx app.Context, e app.Event) {
	ctx.Navigate("/v/"+strconv.Itoa(v.group.Group.Id))
}

func (v *volCheckin) reload(ctx app.Context, e app.Event) {
	ctx.Reload()
}

func (v *volCheckin) Render() app.UI {
	return app.Div().
		Class("py-4 text-center").
		Body(
			app.If(v.group.Group.Attendance.DidAttend == true,
				app.If(v.group.GroupRoleId == rockapi.LeaderRoleId,
					app.Button().
						TabIndex('0').
						Class("w-44 bg-blue-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none mr-2").
						Text("Group Manager").
						OnClick(v.openGroup),
				).Else(
					component.Header().Level(3).Text("You're checked in. Please head to your teams's Huddle location."),
				),
			).Else(
				app.If(v.locStatus == "out",
					app.H4().
						Class("text-base font-normal leading-normal my-1 text-primary").
						Text("Mobile Check-In available when you arrive at Liquid Church."),
					app.Button().
						TabIndex('0').
						Class("w-72 bg-blue-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none mr-2").
						Text("I'm at Liquid. Check-In Now").
						OnClick(v.addAttendance),
				).ElseIf(v.locStatus == "in",
					component.Button().Width("large").Color("secondary").Label("Check-In Now").OnClick(v.addAttendance),
				).ElseIf(v.locStatus == "err",
					app.H4().
						Class("text-base font-normal leading-normal my-1 text-primary").
						Text("Please click on Check In When you get to Liquid."),
					app.Button().
						TabIndex('0').
						Class("w-72 bg-blue-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none mr-2").
						Text("I'm at Liquid. Check-In Now").
						OnClick(v.addAttendance),	
				).Else(
					app.H4().
						Class("text-base font-normal leading-normal my-1 text-primary").
						Text("Determining your location..."),
				),
			),
		)
}