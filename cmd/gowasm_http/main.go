package main

import (
	"net/http"

	httpwasm "github.com/enrichman/gowasm/pkg/http"
)

func main() {
	http.ListenAndServe(":8989", httpwasm.Handler())
}
