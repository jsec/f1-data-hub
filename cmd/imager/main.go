package main

import (
	"context"
	"log"

	"github.com/jsec/f1-data-hub/internal/imager"
)

func main() {
	if err := imager.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
