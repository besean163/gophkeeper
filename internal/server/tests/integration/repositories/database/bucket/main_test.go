package bucket

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/database"
	"github.com/besean163/gophkeeper/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbname     = "gophkeeper_test"
	dbuser     = "gophkeeper_test"
	dbpassword = "gophkeeper_test"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	db = setupTestDB()

	code := m.Run()
	// cleanup()
	os.Exit(code)
}

func setupTestDB() *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbuser, dbpassword, dbhost, dbport, dbname)
	err := database.RunMigrations(dsn)

	if err != nil {
		log.Fatalln("failed migrations", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод логов
		logger.Config{
			LogLevel: logger.Silent, // Уровень логирования: Silent (отключает логи)
		},
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("failed open connect", err)
	}
	return conn
}

// func cleanup() {
// 	query := `
// 	DO $$
// 	BEGIN
// 		EXECUTE (
// 			SELECT string_agg('DROP TABLE IF EXISTS ' || tablename || ' CASCADE;', ' ')
// 			FROM (
// 				SELECT tablename
// 				FROM pg_tables
// 				WHERE schemaname = 'public'
// 			) AS tbl
// 		);
// 	END $$;
// 	`

// 	err := db.Exec(query).Error
// 	if err != nil {
// 		log.Fatalf("Ошибка выполнения SQL-запроса: %v", err)
// 	}

// 	fmt.Println("Все таблицы в схеме public успешно удалены.")
// }

func loadFixtureAccounts(t *testing.T, items []*models.Account) {
	t.Helper()
	for _, item := range items {
		err := db.Create(&item).Error
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
