package http

import (
	"fmt"
	"net/http"

	"github.com/enrichman/gowasm/internal/log"
	"github.com/enrichman/gowasm/internal/view"
	"github.com/gorilla/mux"
)

var (
	Version = "0.0.0-dev"
)

func Router(logger log.Logger, r *mux.Router) {
	if r == nil {
		r = mux.NewRouter()
	}

	r.HandleFunc("/hello", HelloHandler(logger))
	r.HandleFunc("/version", VersionHandler(logger))
}

func HelloHandler(logger log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Log("helloHandler")

		err := req.ParseForm()
		if err != nil {
			logger.Error(err, "req.ParseForm")
		}
		name := req.Form["name"][0]
		logger.Log("helloHandler name: " + name)

		component := view.HiButton(name)
		err = component.Render(req.Context(), res)
		if err != nil {
			logger.Error(err, "component.Render")
		}
	}
}

func VersionHandler(logger log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Log("VersionHandler")
		fmt.Fprint(res, Version)
	}
}
