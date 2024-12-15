package database

import (
	"log"
	"os"
)

func InitializeDatabase(dbPath string) error {
	// Проверяем, существует ли база данных
	if _, err := os.Stat(dbPath); err == nil {
		log.Println("База данных уже существует.")
		return nil
	}

	// Создаем пустую базу данных
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()

	log.Println("Пустая база данных создана.")
	return nil
}
