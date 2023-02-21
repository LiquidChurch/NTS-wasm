package component

import (
	"log"
	"sync"
	"strings"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type QrScanner struct {
	app.Compo

	initScanner  sync.Once
	scanner      app.Value
	onScanSuccess sync.Once
}

var scannerConfig = map[string]interface{}{
	"fps": 3,
}

var cameraSelect = map[string]interface{}{
	"facingMode":"environment",
}

func NewQrScanner() *QrScanner {
	return &QrScanner{}
}

func (q *QrScanner) OnMount(ctx app.Context) {
	q.loadQRScanner()
	q.startQRScanner(ctx)
}

func (q *QrScanner) OnScan(app.Value) {
	log.Println("On Scan Success")
}

func (q *QrScanner) loadQRScanner() {
	log.Println("Load QR Scanner")
	q.initScanner.Do(func() {
		q.scanner = app.Window().
		Get("Html5Qrcode").
		New("qr-reader")
	})
}

func (q *QrScanner) startQRScanner(ctx app.Context) {
	log.Println("Start QR Scanner")
	onScanSuccess := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		log.Println("Scan Success")
		
		scannedArray := strings.Split(args[0].String(),"/")
		if len(scannedArray) == 5 {
			ctx.NewActionWithValue("readAppKey", scannedArray[4])
			q.stopQRScanner(ctx)
		}
		return nil
	})

	q.scanner.Call("start", 
					cameraSelect,
					scannerConfig,
					onScanSuccess,
				)
}

func (q *QrScanner) stopQRScanner(ctx app.Context) {
	log.Println("Stop QR Scanner")
	q.scanner.Call("stop")
}

func (q *QrScanner) Render() app.UI {
	return app.Div().
		Class("qrscanner").
		Body(
			app.Div().
				ID("qr-reader").
				Class("w-66vw max-w-lg h-50vw max-h-96"),
		)
}