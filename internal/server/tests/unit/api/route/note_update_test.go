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
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/models"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNoteUpdateRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	user := models.User{Login: "user", Password: "password"}

	bucketService := mock.NewMockBucketService(ctrl)

	debs := dependencies.Dependencies{
		Logger:        defaultlogger.NewDefaultLogger(),
		BucketService: bucketService,
	}

	handler := route.NoteUpdateRoute(debs)

	tests := []struct {
		name            string
		method          string
		exceptCode      int
		requestBody     string
		responseHeaders map[string]string
		responseBody    string
		user            *models.User
		mockSetup       func()
	}{
		{
			name:        "success",
			method:      http.MethodPut,
			exceptCode:  http.StatusOK,
			requestBody: `{"uuid":"test_uuid","name":"test_name","content":"test_content"}`,
			user:        &user,
			mockSetup: func() {
				bucketService.EXPECT().UpdateNote(user, &models.Note{UUID: "test_uuid", Name: "test_name", Content: "test_content"}).Return(nil).Times(1)
			},
		},
		{
			name:         "without_uuid",
			method:       http.MethodPut,
			exceptCode:   http.StatusBadRequest,
			requestBody:  `{"name":"test_name","content":"test_content"}`,
			responseBody: `{"error":{"code":400,"description":"empty uuid"}}`,
			user:         &user,
			mockSetup:    func() {},
		},
		{
			name:         "not_found_uuid",
			method:       http.MethodPut,
			exceptCode:   http.StatusBadRequest,
			requestBody:  `{"uuid":"not_found_uuid","name":"test_name","content":"test_content"}`,
			responseBody: `{"error":{"code":400,"description":"not found by uuid"}}`,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().UpdateNote(user, &models.Note{UUID: "not_found_uuid", Name: "test_name", Content: "test_content"}).Return(apierrors.ErrorNotFoundByUUID).Times(1)
			},
		},
		{
			name:         "invalid_json",
			method:       http.MethodPut,
			requestBody:  `{`,
			exceptCode:   http.StatusBadRequest,
			responseBody: `{"error":{"code":400,"description":"invalid JSON data"}}`,
			user:         &user,
			mockSetup:    func() {},
		},
		{
			name:   "not_auth",
			method: http.MethodPut,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":401,"description":"not authorized"}}`,
			exceptCode:   http.StatusUnauthorized,
			mockSetup:    func() {},
		},
		{
			name:        "unknown_error",
			method:      http.MethodPut,
			requestBody: `{"uuid":"uuid","name":"test_name","login":"test_login","password":"test_password"}`,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":400,"description":"test"}}`,
			exceptCode:   http.StatusBadRequest,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().UpdateNote(gomock.Any(), gomock.Any()).Return(errors.New("test")).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
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
