package http

import (
	"net/http"

	"github.com/enrichman/gowasm/internal/log"
	"github.com/enrichman/gowasm/view"
	"github.com/gorilla/mux"
)

func Router(logger log.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler(logger))
	return r
}

func helloHandler(logger log.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Log("helloHandler")

		err := req.ParseForm()
		if err != nil {
			logger.Error(err, "req.ParseForm")
		}
		name := req.Form["name"][0]
		logger.Log("helloHandler name: " + name)

		component := view.HelloResult(name, false)
		err = component.Render(req.Context(), res)
		if err != nil {
			logger.Error(err, "component.Render")
		}
	}
}
