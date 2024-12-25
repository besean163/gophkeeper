package database

import (
	"os"

	"github.com/besean163/gophkeeper/internal/logger"
)

func InitializeDatabase(dbPath string) error {
	// Проверяем, существует ли база данных
	if _, err := os.Stat(dbPath); err == nil {
		logger.Debug("База данных уже существует.")
		return nil
	}

	// Создаем пустую базу данных
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()

	logger.Debug("Пустая база данных создана.")
	return nil
}
