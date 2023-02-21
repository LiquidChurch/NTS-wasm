package component

import (
	"log"
	"math"
	"sync"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GeoLocation struct {
	app.Compo

	initLocation  sync.Once
	navigator     app.Value
	onSuccess 	  sync.Once
}

func NewGeoLocation() *GeoLocation {
	return &GeoLocation{}
}

func (g *GeoLocation) OnMount(ctx app.Context) {
	ctx.SetState("geoLocation", "acq")
	onScanSuccess := app.FuncOf(func(this app.Value, args []app.Value) interface{} {

		//Liquid Church
		homeLatitude := 40.85
		homeLongitude := -74.42

		latitude := args[0].JSValue().Get("coords").Get("latitude").Float()
		longitude := args[0].JSValue().Get("coords").Get("longitude").Float()
		
		log.Println(math.Floor(latitude*100)/100, math.Round(longitude*100)/100)
		if (math.Floor(latitude*100)/100 == homeLatitude && math.Round(longitude*100)/100 == homeLongitude) {
			ctx.SetState("geoLocation", "in")
		} else {
			ctx.SetState("geoLocation", "out")
		}
		ctx.NewAction("geoLocUpdate")
		return nil
	})

	onScanError := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		errorCode := args[0].JSValue().Get("code").Int()
		log.Println("Location error", errorCode)
		if errorCode != 0 {
			ctx.SetState("geoLocation", "err")
		}
		ctx.NewAction("geoLocUpdate")
		return nil
	})
	log.Println("mount geo location")
	//g.loadGeoLocation()
	//log.Println(g.navigator)
	app.Window().Get("navigator").Get("geolocation").Call("getCurrentPosition", onScanSuccess, onScanError)
}

func (g *GeoLocation) loadGeoLocation() {	
	g.initLocation.Do(func() {
		g.navigator = app.Window().
		Get("navigator").
		Get("geolocation")
	})

}

func (g *GeoLocation) Render() app.UI {
	return app.Div()
}