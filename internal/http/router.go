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
		logger.Info("helloHandler")

		err := req.ParseForm()
		if err != nil {
			logger.Error("req.ParseForm", "error", err.Error())
		}
		name := req.Form["name"][0]
		logger.Info("helloHandler name: " + name)

		component := view.HiButton(name)
		err = component.Render(req.Context(), res)
		if err != nil {
			logger.Error("component.Render", "error", err.Error())
		}
	}
}

func VersionHandler(logger log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("VersionHandler")
		fmt.Fprint(res, Version)
	}
}
