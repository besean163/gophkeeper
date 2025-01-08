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

	// if err := initDatabase(); err != nil {
	// 	log.Fatalf("Ошибка при инифиализации базы: %v", err)
	// }

	// core.Init()
	// p := app.NewProgram()

	// _, err := p.Run()
	// if err != nil {
	// 	log.Fatalf("Ошибка при запуске программы: %v", err)
	// }
}

// func initDatabase() error {
// 	dbPath := "data.db"
// 	if err := database.InitializeDatabase(dbPath); err != nil {
// 		return err
// 	}
// 	if err := database.RunMigrations(dbPath); err != nil {
// 		return err
// 	}
// 	return nil
// }
