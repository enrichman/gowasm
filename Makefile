build:
	rm -rf dist || true
	mkdir dist
	templ generate
	go run cmd/generate/main.go > dist/index.html
	cp sw.js dist/
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	go run cmd/server/main.go
