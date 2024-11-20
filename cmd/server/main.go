package main

import (
	"context"
	"log"

	"github.com/jsec/f1-data-hub/internal/server"
)

func main() {
	if err := server.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
