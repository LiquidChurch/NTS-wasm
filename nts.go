package main

import (
	"nts/page"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &page.Index{})
	app.Route("/teamlead", &page.TeamLead{})
	app.RouteWithRegexp("/g/*", &page.Guest{})
	app.RouteWithRegexp("/v/*", &page.Group{})
	app.RouteWithRegexp("/*", &page.Lookup{})

	app.RunWhenOnBrowser()

	_ = app.GenerateStaticWebsite("../gae/pwa/", &app.Handler{
		Name:        "Night to Shine",
		Description: "Liquid Church Night to Shine App.",
		Version: 	 "v0.2.15",
		Icon: app.Icon{
			Default:    "/web/asset/liquid-192.png",
			AppleTouch: "/web/asset/apple-icon.png",
		},
		Resources:   app.LocalDir("/"),
		Scripts: []string{
			"https://unpkg.com/html5-qrcode",
		},
		Styles: []string{
			"/web/styles/tailwind.css",
		},
	})
}