//go:build !js

package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir(`dist`))
	http.ListenAndServe(`:8085`, fileServer)
}
