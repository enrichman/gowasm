package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	return r
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	in := req.Form["input"][0]
	fmt.Fprintf(res, `<p id="result">%s (hello changed %s)</p>`, in, req.URL.Path)
}
