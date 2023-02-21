package element

import (
	"log"
	"nts/component"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func GuestInfo() *guestInfo {
	return &guestInfo{}
}

type guestInfo struct {
	app.Compo

	guestAtt 	rockapi.GuestAtt
	infoSet		bool
}

func (g *guestInfo) OnMount(ctx app.Context) {
	g.infoSet = false
	ctx.Handle("guest", g.setInfo)
}

func (g *guestInfo) setInfo(ctx app.Context, a app.Action) {
	g.guestAtt = interface{}(a.Value).(rockapi.GuestAtt)
	g.infoSet = true
}

func (g *guestInfo) Render() app.UI {
	var	infoSet bool
	if g.infoSet {
		log.Println("Guest Info Loaded")
		infoSet = true
	} else {
		log.Println("Guest Info Not Loaded")
		return app.Div().Body()
	}

	return app.Div().
		Class("").
		Body(
			app.If(infoSet == true,
				app.If(g.guestAtt.MedicalPager.Value != "",
					component.InfoCard().Label("Pager:").
						Content(
							[]app.UI{
								component.Header().Level(2).Text(g.guestAtt.MedicalPager.Value),
							}),
				),
				app.If(g.guestAtt.MedicationFlag.Value == "Yes",
					component.InfoCard().Label("Medication").
						Content(
							[]app.UI{
							component.Header().Level(2).Text(g.guestAtt.MedicationNameTime.Value),
							app.Div().
								Class("flex flex-row flex-wrap justify-between").
								Body(
									component.Header().Level(2).Text(g.guestAtt.MedicationAdminister.Value),
									component.Header().Level(3).Text(g.guestAtt.MedicationAdminPhone.Value),
								),
						}),
				),
				component.InfoCard().Label("Emergency Contact").
					Content(
					[]app.UI{
						app.Div().
						Class("flex flex-row flex-wrap justify-between").
						Body(
							component.Header().Level(2).Text(g.guestAtt.EmergencyContactName.Value),
							component.Header().Level(3).Text(g.guestAtt.EmergencyContactRelat.Value),
						),
					app.Div().
						Class("flex flex-row flex-wrap justify-between").
						Body(
							component.Header().Level(2).Text(g.guestAtt.EmergencyContactPhone.Value),
							component.Header().Level(3).Text(g.guestAtt.EmergencyContactEmail.Value),
						),
					}),
				app.If(g.guestAtt.GroupHomeFlag.Value == "Yes",
					component.InfoCard().Label("Group Home").
						Content(
							[]app.UI{
							component.Header().Level(2).Text(g.guestAtt.GroupHome.Value),
							app.Div().
								Class("flex flex-row flex-wrap justify-between").
								Body(
									app.H2().
										Class("text-2xl font-normal flex-none leading-normal my-1 text-blue-800").
										Text(g.guestAtt.GroupHomeStaffName.Value),
									app.If(g.guestAtt.GroupHomeStaffFlag.Value == "Yes",
										app.Label().Class("flex-auto px-2 place-self-center").Text("(on-site)"),
										),
									component.Header().Level(3).Text(g.guestAtt.GroupHomeStaffPhone.Value),
								),
						}),
					),
				component.Header().Level(3).Text("Supervision Level"),
				component.Header().Level(2).Text(g.guestAtt.SupervisionLevel.Value),
				component.Header().Level(3).Text("Communication Level"),
				component.Header().Level(2).Text(g.guestAtt.CommunicationLevel.Value),
				component.Header().Level(3).Text("Bathroom Assistance"),
				component.Header().Level(2).Text(g.guestAtt.BathroomAssistance.Value),
				component.Header().Level(3).Text("Special Needs"),
				component.Header().Level(2).Text(g.guestAtt.SpecialNeedDescription.Value),
				component.Header().Level(3).Text("Food Option"),
				component.Header().Level(2).Text(g.guestAtt.FoodOption.Value),
				component.Header().Level(3).Text("Limo"),
				component.Header().Level(2).Text(g.guestAtt.LimoFlag.Value),				
				component.Header().Level(3).Text("Additional Information"),
				component.Header().Level(2).Text(g.guestAtt.AdditionalInfo.Value),
			),
		)
	}