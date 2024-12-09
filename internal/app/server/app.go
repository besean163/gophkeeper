package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/besean163/gophkeeper/internal/auth"
	"github.com/besean163/gophkeeper/internal/repository"
	"github.com/besean163/gophkeeper/internal/routing"
	"github.com/besean163/gophkeeper/internal/services"
	"golang.org/x/sync/errgroup"
)

type App struct {
	ctx    context.Context
	config *Config
	server *http.Server
	auth.AuthService
}

func NewApp() (*App, error) {
	var err error
	config, err := ReadConfig()
	if err != nil {
		return nil, err
	}

	app := &App{
		ctx:    context.Background(),
		config: config,
	}

	err = app.initAuthService()
	if err != nil {
		return nil, err
	}

	err = app.initServer()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) initAuthService() error {
	r := repository.NewUserRepository()
	s := services.NewAuthService(r)
	app.AuthService = s
	return nil
}

func (app *App) initServer() error {
	handler := routing.NewHandler(app.AuthService)
	app.server = NewServer(app.config.Host, handler)
	return nil
}

func (app *App) Run() error {
	app.RunGracefulStop()
	return app.StartServer()
}

func (app *App) StartServer() error {
	errGroup := errgroup.Group{}
	errGroup.Go(GetShutdownServerRoutine(app.ctx, app.server))
	errGroup.Go(GetRunServerRoutine(app.server))

	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (app *App) RunGracefulStop() {
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
