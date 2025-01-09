// Package app пакет предоставляет абстракцию работы сервера REST API
package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/besean163/gophkeeper/internal/logger"
	zaplogger "github.com/besean163/gophkeeper/internal/logger/zap_logger"
	"github.com/besean163/gophkeeper/internal/server/api"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/user"
	"github.com/besean163/gophkeeper/internal/server/services/auth"
	bucketservice "github.com/besean163/gophkeeper/internal/server/services/bucket"
	apitoken "github.com/besean163/gophkeeper/internal/utils/api_token"
	jwttoken "github.com/besean163/gophkeeper/internal/utils/api_token/jwt_token"
	bcryptencrypter "github.com/besean163/gophkeeper/internal/utils/password_encrypter/bcrypt"
	standarttimecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller/standart_time_controller"
	standartuuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller/standart_uuid_controller"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App структура приложения
type App struct {
	logger logger.Logger
	ctx    context.Context
	config *Config
	server *http.Server
	interfaces.AuthService
	interfaces.BucketService
	apitoken.Tokener
}

// NewApp создание структуры приложения
func NewApp() (*App, error) {
	var err error
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	app := &App{
		ctx:    context.Background(),
		config: config,
	}

	err = app.initLogger()
	if err != nil {
		return nil, err
	}

	app.initTokener()

	err = app.initAuthService()
	if err != nil {
		return nil, err
	}

	err = app.initBucketService()
	if err != nil {
		return nil, err
	}

	err = app.initServer()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) initLogger() error {
	logger, err := zaplogger.NewLogger()
	if err != nil {
		return err
	}

	app.logger = logger
	return nil
}

func (app *App) initTokener() {
	app.Tokener = jwttoken.NewTokener(app.config.Secret)
}

func (app *App) initAuthService() error {
	timeController := standarttimecontroller.NewTimeController()
	uuidController := standartuuidcontroller.NewUUIDController()
	encrypter := bcryptencrypter.NewEncrypter()
	db, err := getDB()
	if err != nil {
		return err
	}
	r := user.NewRepository(db)

	options := auth.ServiceOptions{
		Repository:     r,
		Encrypter:      encrypter,
		Tokener:        app.Tokener,
		TimeController: timeController,
		UUIDController: uuidController,
	}

	s := auth.NewService(options)
	app.AuthService = s
	return nil
}

func (app *App) initBucketService() error {
	timeController := standarttimecontroller.NewTimeController()
	uuidController := standartuuidcontroller.NewUUIDController()
	db, err := getDB()
	if err != nil {
		return err
	}
	r := bucket.NewRepository(db, uuidController)

	options := bucketservice.ServiceOptions{
		Repository:     r,
		TimeController: timeController,
		UUIDController: uuidController,
	}

	s := bucketservice.NewService(options)
	app.BucketService = s
	return nil
}

func (app *App) initServer() error {
	dependencies := dependencies.Dependencies{
		Logger:        app.logger,
		Tokener:       app.Tokener,
		AuthService:   app.AuthService,
		BucketService: app.BucketService,
	}

	handler := api.NewHandler(dependencies)
	app.server = newServer(app.config.Host, handler)
	return nil
}

// Run запуск приложения
func (app *App) Run() error {
	app.runGracefulStop()
	return app.startServer()
}

func (app *App) startServer() error {
	errGroup := errgroup.Group{}
	errGroup.Go(getShutdownServerRoutine(app.ctx, app.server))
	errGroup.Go(getRunServerRoutine(app.server))

	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (app *App) runGracefulStop() {
	ctx, cancel := context.WithCancel(app.ctx)
	app.ctx = ctx
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		log.Println("graceful stop")
		cancel()
	}()
}

func getDB() (*gorm.DB, error) {
	dsn := "postgres://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
