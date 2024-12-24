package main

import (
	"log"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/database"
	app "github.com/besean163/gophkeeper/internal/client/tui"
	"github.com/besean163/gophkeeper/internal/logger"
)

func main() {

	logger.Debug("here")
	if err := initDatabase(); err != nil {
		log.Fatalf("Ошибка при инифиализации базы: %v", err)
		logger.Debug("her1")
	}

	core.Init()
	p := app.NewProgram()

	_, err := p.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске программы: %v", err)
	}
}

func initDatabase() error {
	dbPath := "data.db"
	if err := database.InitializeDatabase(dbPath); err != nil {
		return err
	}
	if err := database.RunMigrations(dbPath); err != nil {
		return err
	}
	return nil
}
