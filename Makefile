build:
	templ generate
	go run cmd/generate/main.go > dist/index.html
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	go run cmd/server/main.go
