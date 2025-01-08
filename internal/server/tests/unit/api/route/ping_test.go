package route

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	handler := route.PingRoute()
	request, _ := http.NewRequest(http.MethodGet, "", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)

	response := rr.Result()
	defer response.Body.Close()

	// проверяем код ответа
	assert.Equal(t, http.StatusOK, response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "pong", string(responseBody))
}
