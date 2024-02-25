package main

import (
	"context"
	"os"

	"github.com/enrichman/gowasm/internal/view"
)

func main() {
	indexFile, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}

	index := view.Index(true)
	err = index.Render(context.Background(), indexFile)
	if err != nil {
		panic(err)
	}
}
