package user

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/database"
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
	log.Println("user in")
	db = setupTestDB()

	code := m.Run()
	// cleanup()
	log.Println("user out")
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
