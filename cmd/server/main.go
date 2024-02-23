//go:build !js

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}
	fileServer := http.FileServer(http.Dir(`dist`))

	fmt.Println("listening on http://localhost:" + port)
	err := http.ListenAndServe(`:`+port, fileServer)
	if err != nil {
		log.Fatal(err)
	}
}
