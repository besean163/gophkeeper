package main

import (
	"log"

	"github.com/besean163/gophkeeper/internal/app/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	app, err := server.NewApp()
	if err != nil {
		return err
	}

	return app.Run()
}
