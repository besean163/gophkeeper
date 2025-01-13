// Package app пакет предоставляет абстракцию работы приложения TUI
package app

import (
	"log"
	"os"

	"github.com/besean163/gophkeeper/internal/client/core"
	store "github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/database"
	dbinit "github.com/besean163/gophkeeper/internal/client/database"
	"github.com/besean163/gophkeeper/internal/client/interfaces"
	tui "github.com/besean163/gophkeeper/internal/client/tui"
	"github.com/besean163/gophkeeper/internal/logger"
	zaplogger "github.com/besean163/gophkeeper/internal/logger/zap_logger"
	"github.com/besean163/gophkeeper/internal/server/api/client"
	restyclient "github.com/besean163/gophkeeper/internal/server/api/client/http_client/resty_client"
	bcryptencrypter "github.com/besean163/gophkeeper/internal/utils/password_encrypter/bcrypt"
	standarttimecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller/standart_time_controller"
	standartuuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller/standart_uuid_controller"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	dblogger "gorm.io/gorm/logger"
)

var dbpath = "data.db"

// App структура приложения
type App struct {
	logger logger.Logger
	core   interfaces.Core
	ui     tui.App
}

// NewApp создание структуры приложения
func NewApp() (*App, error) {
	var err error
	app := &App{}

	err = app.setLogger()
	if err != nil {
		return nil, err
	}

	err = app.setCore()
	if err != nil {
		return nil, err
	}

	err = app.initDatabase()
	if err != nil {
		return nil, err
	}

	app.setUI()

	return app, nil
}

// Run запуск приложения
func (app App) Run() error {
	return app.ui.Run()
}

func (app *App) setLogger() error {
	logger, err := zaplogger.NewLogger()
	if err != nil {
		return err
	}

	app.logger = logger
	return nil
}

func (app *App) setUI() {
	app.ui = tui.NewApp(app.core, app.logger)
}

func (app *App) setCore() error {

	db, err := getDB()
	if err != nil {
		return err
	}

	encrypter := bcryptencrypter.NewEncrypter()
	databaseRepository := store.NewRepository(db, standartuuidcontroller.NewUUIDController())

	options := database.ServiceOptions{
		Repository:     databaseRepository,
		Encrypter:      encrypter,
		Logger:         app.logger,
		TimeController: standarttimecontroller.NewTimeController(),
		UUIDController: standartuuidcontroller.NewUUIDController(),
	}

	databaseService := database.NewService(options)

	apiClient := client.NewClient("http://localhost:8080", restyclient.NewHTTPClient(), app.logger)

	apiServiceOptions := api.ServiceOptions{
		DataService:    databaseService,
		ApiClient:      apiClient,
		Encrypter:      encrypter,
		TimeController: standarttimecontroller.NewTimeController(),
		Logger:         app.logger,
	}

	apiService := api.NewService(apiServiceOptions)

	core := core.NewCore(apiService, app.logger)
	app.core = &core

	return nil
}

func getDB() (*gorm.DB, error) {
	newLogger := dblogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Вывод логов
		dblogger.Config{
			LogLevel: dblogger.Silent, // Уровень логирования: Silent (отключает логи)
		},
	)
	return gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		Logger: newLogger,
	})
}

func (app *App) initDatabase() error {
	var err error
	err = dbinit.InitializeDatabase(dbpath, app.logger)
	if err != nil {
		app.logger.Error("database init", logger.NewField("error", err.Error()))
		return err
	}

	err = dbinit.RunMigrations(dbpath, app.logger)
	if err != nil {
		app.logger.Error("migration run", logger.NewField("error", err.Error()))
		return err
	}
	return nil
}
