package main

import (
	"log"

	"github.com/nnaxim/train-radar/internal/app"
)

func main() {
	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	application.Run()
}
