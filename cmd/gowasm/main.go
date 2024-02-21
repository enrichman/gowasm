package main

import (
	"github.com/enrichman/gowasm/pkg/http"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	handler := http.Handler()
	wasmhttp.Serve(handler)
	select {}
}
