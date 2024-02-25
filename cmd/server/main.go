//go:build !js

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	gowasmhttp "github.com/enrichman/gowasm/internal/http"
	"github.com/enrichman/gowasm/internal/log"
	"github.com/enrichman/gowasm/internal/view"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewStdLogger()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api/").Subrouter()
	gowasmhttp.Router(logger, apiRouter)

	r.Handle("/", templ.Handler(view.Index(false)))

	staticCSSFileServer := http.FileServer(http.Dir(`dist/css`))
	cssHandler := http.StripPrefix("/css/", staticCSSFileServer)
	r.PathPrefix("/css/").Handler(cssHandler)

	fmt.Println("listening on http://localhost:" + port)
	fmt.Println(gowasmhttp.Version)
	err := http.ListenAndServe(`localhost:`+port, r)
	if err != nil {
		logger.Error(err, "")
	}
}
