
build: templ css
	cp assets/* dist
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	go run cmd/server/main.go

templ:
	templ generate
	go run cmd/generate/main.go

css:
	tailwindcss -i ./src/input.css -o ./dist/output.css

css-watch:
	tailwindcss -i ./src/input.css -o ./dist/output.css --watch

clean:
	rm -rf dist || true
	mkdir dist
