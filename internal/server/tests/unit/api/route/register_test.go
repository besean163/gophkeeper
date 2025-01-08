package route

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoute(t *testing.T) {

	ctrl := gomock.NewController(t)
	authService := mock.NewMockAuthService(ctrl)
	authService.EXPECT().RegisterUser("user", "password").Return("user_token", nil).Times(1)
	authService.EXPECT().RegisterUser("exist_user", "exist_password").Return("", apierrors.ErrorUserExist).Times(1)
	authService.EXPECT().RegisterUser("create_user", "create_password").Return("", errors.New("test")).Times(1)

	debs := dependencies.Dependencies{
		Logger:      defaultlogger.NewDefaultLogger(),
		AuthService: authService,
	}

	handler := route.RegisterRoute(debs)

	tests := []struct {
		name            string
		method          string
		exceptCode      int
		requestBody     string
		responseHeaders map[string]string
		responseBody    string
	}{
		{
			name:        "success",
			method:      http.MethodPost,
			requestBody: `{"login":"user","password":"password"}`,
			exceptCode:  http.StatusOK,
			responseHeaders: map[string]string{
				"Content-type": "application/json",
			},
			responseBody: `{"token":"user_token"}`,
		},
		{
			name:        "empty_login",
			method:      http.MethodPost,
			requestBody: `{"login":"","password":"password"}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"login empty"}}`,
			exceptCode:   http.StatusBadRequest,
		},
		{
			name:        "empty_password",
			method:      http.MethodPost,
			requestBody: `{"login":"login","password":""}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"password empty"}}`,
			exceptCode:   http.StatusBadRequest,
		},
		{
			name:        "exist_user",
			method:      http.MethodPost,
			requestBody: `{"login":"exist_user","password":"exist_password"}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"user already exist"}}`,
			exceptCode:   http.StatusBadRequest,
		},
		{
			name:        "invalid_json",
			method:      http.MethodPost,
			requestBody: `{"login":"test_user","password":"test_passwo`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"invalid JSON data"}}`,
			exceptCode:   http.StatusBadRequest,
		},
		{
			name:        "unknown_error",
			method:      http.MethodPost,
			requestBody: `{"login":"create_user","password":"create_password"}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":500,"description":"unknown internal error"}}`,
			exceptCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := strings.NewReader(test.requestBody)
			request, _ := http.NewRequest(test.method, "/api/user/register", requestBody)
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
