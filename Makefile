
build: templ css
	cp static/sw.js dist/sw.js
	go run cmd/generate/main.go
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	air

templ:
	templ generate

templ-watch:
	templ generate --watch

css:
	tailwindcss -i ./static/css/input.css -o ./dist/css/output.css --minify

css-watch:
	tailwindcss -i ./static/css/input.css -o ./dist/css/output.css --watch

clean:
	rm -rf dist || true
	mkdir dist

serve:
	docker run --rm -p8082:80 -it -v $(shell pwd)/dist:/usr/share/nginx/html nginx
