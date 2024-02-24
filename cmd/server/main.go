//go:build !js

package main

import (
	"fmt"
	"net/http"
	"os"

	gowasmhttp "github.com/enrichman/gowasm/http"
	"github.com/enrichman/gowasm/internal/log"
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

	fileServer := http.FileServer(http.Dir(`dist`))
	r.PathPrefix("/").Handler(fileServer)

	fmt.Println("listening on http://localhost:" + port)
	err := http.ListenAndServe(`localhost:`+port, r)
	if err != nil {
		logger.Error(err, "")
	}
}
