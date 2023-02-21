package page

import (
	"log"
	"strings"
	"nts/rockapi"
	"nts/component"
	"nts/element"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Guest struct {
	app.Compo
	appKey 		string
	isEdit		bool
	guest		*rockapi.GuestDetail
	guestId		rockapi.GuestId
}


func (g *Guest) OnMount(ctx app.Context) {
	g.updateGuest(ctx)
}

func (g *Guest) OnAppUpdate(ctx app.Context) {
	updateAvailable := ctx.AppUpdateAvailable() 

	if updateAvailable == true {
		ctx.Reload()
	}	
}


func (g *Guest) updateGuest(ctx app.Context) {
	g.appKey = strings.TrimPrefix(ctx.Page().URL().Path, "/g/")
	log.Println("AppKey:", g.appKey)
	g.guest = rockapi.LoadGuestByAppKey(ctx, g.appKey)

	if (g.guest != nil) {
		ctx.NewActionWithValue("person", g.guest.Person)
		ctx.NewActionWithValue("guest", g.guest.Attributes)
		g.guestId.GuestId = g.guest.Id
		g.guestId.PersonId = g.guest.Person.Id
		g.guestId.PersonAliasId = g.guest.PersonAliasId
		ctx.NewActionWithValue("guestCheckin", g.guestId)
	}
}

func (g *Guest) Render() app.UI {
	return app.Main().
	Class("").
	Body(
		component.TopBar().
			Right(
				[]app.UI{
					element.GuestCheckin(),
				}),
		app.Div().
			Class("p-4").
			Body(
				app.Div().
					Class("relative flex flex-col min-w-0 break-words border-gray-200 border-2 shadow-md rounded-2xl m-2 p-2 bg-clip-border").
					Body(
						element.Person(),
						element.GuestInfo(),
					),
				),
	)
}




/*
func OpenGuest() *Guest {
	return &Guest{}
}

func (g *Guest) goHome(ctx app.Context, e app.Event) {
	ctx.Navigate("/")
}

func (g *Guest) Render() app.UI {
	return app.Main().
		Class("guest p-8").
		Body(
			component.TopBar(),
			app.H1().
			Text("Guest Lookup"),
			component.Card(),
			component.ButtonGroup().ButtonList(
				[]app.UI{
					component.Button().Width("large").Color("blue-200").Label("Home").OnClick(g.goHome),
				}),
		)
}
*/