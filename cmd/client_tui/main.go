package main

import (
	"log"

	"github.com/besean163/gophkeeper/internal/client/app"
)

func main() {
	app, err := app.NewApp()

	if err != nil {
		log.Fatalf("Ошибка при инициализации приложения: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Ошибка при запуске программы: %v", err)
	}
}
