package modal

import (
	"log"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type checkinModal struct {
	app.Compo

	title string
	member 	rockapi.GroupMemberDetail
}

func (m *checkinModal) OnMount(ctx app.Context) {
	log.Println("Checkin Modal Mount")
	ctx.ObserveState("checkinModal").Value(&m.member)
	log.Println(m.member)
}

func (m *checkinModal) checkIn(ctx app.Context, e app.Event) {
	log.Println("button clicked")
	if m.member.Attendance.DidAttend == false {
		log.Println("Add Attendance Triggered")
		rockapi.AddAttendanceByGroupId(ctx, m.member.PersonId, m.member.GroupId)
	} else {
		log.Println("Delete Attendance Triggered")
		rockapi.DelAttendanceById(ctx, m.member.Attendance.Id)
	}
	ctx.Reload()
}

func (m *checkinModal) Render() app.UI {
	flag := "Not Checked In"
	btn := "Check-In"
	if m.member.Attendance.DidAttend == true {
		flag = "Checked In"
		btn = "Uncheck-In"
	} 
	return &modal{
		Title:        m.member.Person.FullName,
		DisableFocus: false,
		Body: []app.UI {
				app.H3().
					Class("text-xl font-normal leading-normal my-1 pl-2 text-primary").
					Text("Email:"),
				app.H3().
					Class("text-xl font-normal leading-normal my-1 pl-4 text-primary").
					Text(m.member.Person.Email),
				app.H3().
					Class("text-xl font-normal leading-normal my-1 text-primary").
					Text(flag),
				app.If(m.member.Attendance.DidAttend == false,
					app.Button().
						TabIndex('0').
						Class("w-44 mx-auto bg-blue-600 py-2 px-4 shadow-md no-underline rounded-full text-white font-sans font-semibold text-sm border-blue hover:opacity-75 focus:outline focus:outline-focus active:shadow-none").
						Text(btn).
						OnClick(m.checkIn),
				),
			},
	}
}
