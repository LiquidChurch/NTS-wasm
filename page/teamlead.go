package page

import (
	"nts/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type TeamLead struct {
	app.Compo
	path string
}

func (h *TeamLead) OnAppUpdate(ctx app.Context) {
	updateAvailable := ctx.AppUpdateAvailable() 

	if updateAvailable == true {
		ctx.Reload()
	}	
}


func (h *TeamLead) Render() app.UI {
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
						component.Header().Level(1).Text("Team Lead Info"),
						component.Header().Level(3).Text("Documents:"),
						app.A().
							Class("text-xl font-normal leading-normal text-blue-800 pl-4").
							Text("Event Map").
							Href("https://drive.google.com/file/d/1LAE5NPvddEreQ-hEvhVNTeSQRf5bQGk-/view"),
						app.A().
							Class("text-xl font-normal leading-normal text-blue-800 pl-4").
							Text("Prom Run Sheet").
							Href("https://drive.google.com/file/d/19_gpO9S365b7Mm4sYFi5LtbwuCPVN2Jh/view"),	
						app.A().
							Class("text-xl font-normal leading-normal text-blue-800 pl-4").
							Text("Walkie Talkie Channel Guide").
							Href("https://drive.google.com/file/d/1pjYFReZ4czMDH14pfXB8vCVVVJu3hG4w/view"),
						app.A().
							Class("text-xl font-normal leading-normal text-blue-800 pl-4").
							Text("Team Huddle Location").
							Href("https://drive.google.com/file/d/1kYXUfraPCGi1VMXK24urhBUi7G9_dcwj/view"),	
					),
				),
		)
}