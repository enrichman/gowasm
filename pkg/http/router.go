package http

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir(`.`)))
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/api/hello", helloHandler)

	return r
}

//go:embed t1.html
var t1 embed.FS

func helloHandler(res http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.ParseFS(t1, "t1.html"))

	err := templ.ExecuteTemplate(res, "t1.html", map[string]string{
		"Name": "pippo",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//	fmt.Fprint(res, `
	//	<p>{{.Name}}</p>
	//	<button hx-get="api/hello" hx-target="#user-info" hx-swap="innerHTML">Reveal Info</button>
	//
	// `)
}
