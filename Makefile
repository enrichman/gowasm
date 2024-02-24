
build:
	go run cmd/generate/main.go
	cp assets/* dist
	GOOS=js GOARCH=wasm go build -o dist/server.wasm cmd/server/main_js_wasm.go

run:
	go run cmd/server/main.go

templ:
	templ generate

css:
	tailwindcss -i ./src/input.css -o ./dist/output.css

clean:
	rm -rf dist || true
	mkdir dist

serve:
	docker run --rm -p8082:80 -it -v $(shell pwd)/dist:/usr/share/nginx/html nginx