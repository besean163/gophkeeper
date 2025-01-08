package route

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundRoute(t *testing.T) {
	debs := dependencies.Dependencies{
		Logger: defaultlogger.NewDefaultLogger(),
	}

	handler := route.NotFoundRoute(debs)

	tests := []struct {
		name            string
		method          string
		exceptCode      int
		responseHeaders map[string]string
		responseBody    string
	}{
		{
			name:       "success",
			method:     http.MethodGet,
			exceptCode: http.StatusNotFound,
			responseHeaders: map[string]string{
				"Content-type": "application/json",
			},
			responseBody: `{"error":{"code":404,"description":"page not found"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := http.NewRequest(test.method, "", nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, request)

			response := rr.Result()
			defer response.Body.Close()

			// проверяем код ответа
			assert.Equal(t, test.exceptCode, response.StatusCode)

			// проверяем заголовки
			for k, v := range test.responseHeaders {
				assert.Equal(t, v, response.Header.Get(k))
			}

			// проверяем тело ответа
			responseBody, _ := io.ReadAll(response.Body)
			assert.Equal(t, test.responseBody, string(responseBody))
		})

	}

}
