package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
)

func TestHadler(t *testing.T) {
	handler := api.NewHandler(dependencies.Dependencies{
		Logger: defaultlogger.NewDefaultLogger(),
	})

	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)
}
