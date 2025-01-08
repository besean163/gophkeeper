package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	dbinit "github.com/besean163/gophkeeper/internal/client/database"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbpath = "data.db"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	db = setupTestDB()

	code := m.Run()
	cleanup()
	os.Exit(code)
}

func setupTestDB() *gorm.DB {
	var err error
	err = dbinit.InitializeDatabase(dbpath, defaultlogger.NewDefaultLogger())
	if err != nil {
		log.Fatalln("failed init db", err)
	}

	err = dbinit.RunMigrations(dbpath, defaultlogger.NewDefaultLogger())
	if err != nil {
		log.Fatalln("failed migrations", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод логов
		logger.Config{
			LogLevel: logger.Silent, // Уровень логирования: Silent (отключает логи)
		},
	)

	conn, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("failed open connect", err)
	}
	return conn
}

func cleanup() {
	err := os.Remove(dbpath)
	if err != nil {
		log.Fatalln("failed cleanup db", err)
	}
	fmt.Println("database file deleted.")
}
