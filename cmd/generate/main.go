package main

import (
	"context"
	"os"

	"github.com/enrichman/gowasm/view"
)

func main() {
	index := view.Index()
	index.Render(context.Background(), os.Stdout)
}
