//go:build js

package main

import (
	"github.com/enrichman/gowasm/http"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	router := http.Router()
	wasmhttp.Serve(router)
	select {}
}
