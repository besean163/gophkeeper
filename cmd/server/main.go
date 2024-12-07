package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/besean163/gophkeeper/internal/routing"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {

	serverCtx, serverStop := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		log.Println("grace work")
		serverStop()
	}()

	handler := routing.NewHandler()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	errGroup := errgroup.Group{}
	errGroup.Go(
		func() error {
			<-serverCtx.Done()
			log.Println("shutdown work")
			if err := server.Shutdown(serverCtx); err != nil {
				log.Fatalf("server shutdown error: %s", err.Error())
				return err
			}
			log.Println("server stopped")
			return nil
		},
	)

	errGroup.Go(
		func() error {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Println(err.Error())
				return err
			}
			return nil
		},
	)

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}
