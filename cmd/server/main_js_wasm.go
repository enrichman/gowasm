//go:build js

package main

import (
	"github.com/enrichman/gowasm/http"
	"github.com/enrichman/gowasm/internal/log"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	logger := log.NewJSLogger()
	router := http.Router(logger)
	wasmhttp.Serve(router)

	select {}
}
