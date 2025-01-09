package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	dbinit "github.com/besean163/gophkeeper/internal/client/database"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
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

func loadFixtureUsers(t *testing.T, users []*models.User) {
	t.Helper()
	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureUser(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM users").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func loadFixtureNotes(t *testing.T, items []*models.Note) {
	t.Helper()
	for _, item := range items {
		err := db.Create(&item).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureNotes(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM notes").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func loadFixtureCards(t *testing.T, items []*models.Card) {
	t.Helper()
	for _, item := range items {
		err := db.Create(&item).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureCards(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM cards").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func loadFixtureAccounts(t *testing.T, accounts []*models.Account) {
	t.Helper()
	for _, account := range accounts {
		err := db.Create(&account).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureAccounts(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM accounts").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}
