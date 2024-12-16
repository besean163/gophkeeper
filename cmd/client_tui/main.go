package main

import (
	"log"

	app "github.com/besean163/gophkeeper/internal/client/tui"
)

func main() {
	p := app.NewProgram()

	_, err := p.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске программы: %v", err)
	}
}
