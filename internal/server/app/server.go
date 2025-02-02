package app

import (
	"context"
	"log"
	"net/http"
)

func newServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func getRunServerRoutine(server *http.Server) func() error {
	return func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err.Error())
			return err
		}
		return nil
	}
}

func getShutdownServerRoutine(ctx context.Context, server *http.Server) func() error {
	return func() error {
		<-ctx.Done()
		log.Println("shutdown work")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown error: %s", err.Error())
			return err
		}
		log.Println("server stopped")
		return nil
	}
}
