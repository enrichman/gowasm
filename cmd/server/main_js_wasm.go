//go:build js

package main

import (
	"github.com/enrichman/gowasm/internal/http"
	"github.com/enrichman/gowasm/internal/log"
	"github.com/gorilla/mux"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	logger := log.NewJSLogger()
	router := mux.NewRouter()
	http.Router(logger, router)
	wasmhttp.Serve(router)

	select {}
}
