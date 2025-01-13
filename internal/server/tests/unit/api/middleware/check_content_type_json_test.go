package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/api/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCheckContentTypeJSONMiddleware(t *testing.T) {

	handler := middleware.CheckContentTypeJSONMiddleware()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("success")) })

	tests := []struct {
		name         string
		code         int
		responseBody string
		headers      map[string]string
	}{
		{
			name:         "success",
			code:         http.StatusOK,
			responseBody: `success`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name:         "fail",
			code:         http.StatusBadRequest,
			responseBody: `{"error":{"code":400,"description":"expect JSON data"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "", nil)
			for k, v := range test.headers {
				request.Header.Set(k, v)
			}
			rr := httptest.NewRecorder()

			handler(testHandler).ServeHTTP(rr, request)

			response := rr.Result()
			defer response.Body.Close()

			assert.Equal(t, test.code, response.StatusCode)

			responseBody, _ := io.ReadAll(response.Body)
			assert.Equal(t, test.responseBody, string(responseBody))
		})
	}
}
