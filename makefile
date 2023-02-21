all: web app

web:
	tailwind -i ./tailwind.css -o ../gae/pwa/web/styles/tailwind.css
	GOARCH=wasm GOOS=js go build -o ../gae/pwa/web/app.wasm nts.go

app:
	go build
	./nts
