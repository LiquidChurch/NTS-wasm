package page

import (
	"nts/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Index struct {
	app.Compo
	updateAvailable bool
}

func (h *Index) OnMount(ctx app.Context) {
	ctx.Handle("readAppKey", h.appKeyNavigate)
}

func (h *Index) OnAppUpdate(ctx app.Context) {
	updateAvailable := ctx.AppUpdateAvailable() 

	if updateAvailable == true {
		ctx.Reload()
	}	
}


func (h *Index) appKeyNavigate(ctx app.Context, a app.Action) {
	appKey := interface{}(a.Value).(string)
	ctx.Navigate("/g/" + appKey)
}

func (h *Index) Render() app.UI {
	return app.Main().
		Class("").
		Body(
			component.TopBar().Right(
				[]app.UI{
					app.P().Class("text-right pr-4").Text("v0.2.13"),
				}),
			app.Div().
				ID("QRScanner").
				Class("flex flex-col h-80vh justify-center items-center").
				Body(
					app.H1().
						Text("Please scan NTS pass."),
					component.NewQrScanner(),
				),
		)
}