package route

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNotesRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	now := time.Now()
	user := models.User{Login: "user", Password: "password"}
	bucketService := mock.NewMockBucketService(ctrl)

	debs := dependencies.Dependencies{
		Logger:        defaultlogger.NewDefaultLogger(),
		BucketService: bucketService,
	}

	handler := route.NotesRoute(debs)

	tests := []struct {
		name            string
		method          string
		exceptCode      int
		responseHeaders map[string]string
		responseBody    string
		user            *models.User
		mockSetup       func()
	}{
		{
			name:   "success",
			method: http.MethodGet,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: fmt.Sprintf(`{"notes":[{"uuid":"uuid_1","name":"name_1","content":"test_content","created_at":%d,"updated_at":%d}]}`, now.Unix(), now.Unix()),
			exceptCode:   http.StatusOK,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().GetNotes(user).Return([]*models.Note{
					{
						UUID:      "uuid_1",
						Name:      "name_1",
						Content:   "test_content",
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				}, nil).Times(1)
			},
		},
		{
			name:   "not_auth",
			method: http.MethodGet,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":401,"description":"not authorized"}}`,
			exceptCode:   http.StatusUnauthorized,
			mockSetup:    func() {},
		},
		{
			name:   "unknown_error",
			method: http.MethodGet,
			responseHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			responseBody: `{"error":{"code":500,"description":"unknown internal error"}}`,
			exceptCode:   http.StatusInternalServerError,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().GetNotes(user).Return(nil, errors.New("test")).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			request, _ := http.NewRequest(test.method, "", nil)
			if test.user != nil {
				ctx := context.WithValue(context.Background(), entities.RequestUserKey("user"), test.user)
				request, _ = http.NewRequestWithContext(ctx, test.method, "", nil)
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
