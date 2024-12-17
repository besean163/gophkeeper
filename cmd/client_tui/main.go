package main

import (
	"log"

	"github.com/besean163/gophkeeper/internal/client/core"
	app "github.com/besean163/gophkeeper/internal/client/tui"
)

func main() {

	core.Init()
	p := app.NewProgram()

	_, err := p.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске программы: %v", err)
	}
}
