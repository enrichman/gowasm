
build:
	rm -rf dist || true
	mkdir -p dist/js dist/css
	templ generate
	cp assets/* dist
	go run cmd/generate/main.go
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	go run cmd/server/main.go
