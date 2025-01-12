package route

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAccountUpdateRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	user := models.User{Login: "user", Password: "password"}

	bucketService := mock.NewMockBucketService(ctrl)
	bucketService.EXPECT().UpdateAccount(user, &models.Account{UUID: "00000000-0000-0000-0000-000000000000", Name: "test_name", Login: "test_login", Password: "test_password"}).Return(nil).Times(1)
	bucketService.EXPECT().UpdateAccount(user, &models.Account{UUID: "00000000-0000-0000-0000-000000000000", Name: "test_name", Login: "test_login", Password: "test_password"}).Return(apierrors.ErrorNotFoundByUUID).Times(1)
	bucketService.EXPECT().UpdateAccount(gomock.Any(), gomock.Any()).Return(errors.New("test")).Times(1)

	debs := dependencies.Dependencies{
		Logger:        defaultlogger.NewDefaultLogger(),
		BucketService: bucketService,
	}

	handler := route.AccountUpdateRoute(debs)

	tests := []struct {
		name            string
		method          string
		exceptCode      int
		requestBody     string
		responseHeaders map[string]string
		responseBody    string
		user            *models.User
	}{
		{
			name:        "success_put",
			method:      http.MethodPut,
			exceptCode:  http.StatusOK,
			requestBody: `{"uuid":"00000000-0000-0000-0000-000000000000","name":"test_name","login":"test_login","password":"test_password"}`,
			user:        &user,
		},
		{
			name:         "without_uuid",
			method:       http.MethodPut,
			exceptCode:   http.StatusBadRequest,
			requestBody:  `{"name":"test_name","login":"test_login","password":"test_password"}`,
			responseBody: `{"error":{"code":400,"description":"empty uuid"}}`,
			user:         &user,
		},
		{
			name:         "not_found_uuid",
			method:       http.MethodPut,
			exceptCode:   http.StatusBadRequest,
			requestBody:  `{"uuid":"00000000-0000-0000-0000-000000000000","name":"test_name","login":"test_login","password":"test_password"}`,
			responseBody: `{"error":{"code":400,"description":"not found by uuid"}}`,
			user:         &user,
		},
		{
			name:         "invalid_json",
			method:       http.MethodPut,
			requestBody:  `{`,
			exceptCode:   http.StatusBadRequest,
			responseBody: `{"error":{"code":400,"description":"invalid JSON data"}}`,
			user:         &user,
		},
		{
			name:   "not_auth",
			method: http.MethodPut,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":401,"description":"not authorized"}}`,
			exceptCode:   http.StatusUnauthorized,
		},
		{
			name:        "unknown_error",
			method:      http.MethodPut,
			requestBody: `{"uuid":"00000000-0000-0000-0000-000000000000","name":"test_name","login":"test_login","password":"test_password"}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"test"}}`,
			exceptCode:   http.StatusBadRequest,
			user:         &user,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestBody := strings.NewReader(test.requestBody)
			request, _ := http.NewRequest(test.method, "", requestBody)
			if test.user != nil {
				ctx := context.WithValue(context.Background(), entities.RequestUserKey("user"), test.user)
				request, _ = http.NewRequestWithContext(ctx, test.method, "", requestBody)
			}
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
