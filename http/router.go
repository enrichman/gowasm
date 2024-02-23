package http

import (
	"net/http"

	"github.com/enrichman/gowasm/view"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	return r
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	name := req.Form["name"][0]

	component := view.HelloResult(name)
	component.Render(req.Context(), res)
}
