package middleware

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/middleware"
	"github.com/besean163/gophkeeper/internal/server/models"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)

	authService := mock.NewMockAuthService(ctrl)
	tokener := utilmock.NewMockTokener(ctrl)

	debs := dependencies.Dependencies{
		Logger:      defaultlogger.NewDefaultLogger(),
		AuthService: authService,
		Tokener:     tokener,
	}

	handler := middleware.AuthMiddleware(debs)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("success")) })

	tests := []struct {
		name         string
		token        string
		code         int
		responseBody string
		mockSetup    func()
	}{
		{
			name:         "success",
			token:        "correct_token",
			code:         http.StatusOK,
			responseBody: `success`,
			mockSetup: func() {
				tokener.EXPECT().GetUserId("correct_token").Return(1, nil).Times(1)
				authService.EXPECT().GetUser(1).Return(&models.User{ID: 1, Login: "login"}, nil).Times(1)
			},
		},
		{
			name:         "fail_get_user_id",
			code:         http.StatusInternalServerError,
			responseBody: `{"error":{"code":500,"description":"test_error"}}`,
			mockSetup: func() {
				tokener.EXPECT().GetUserId(gomock.Any()).Return(0, errors.New("test_error")).Times(1)
			},
		},
		{
			name:         "fail_get_user",
			code:         http.StatusInternalServerError,
			responseBody: `{"error":{"code":500,"description":"test_error"}}`,
			mockSetup: func() {
				tokener.EXPECT().GetUserId(gomock.Any()).Return(1, nil).Times(1)
				authService.EXPECT().GetUser(1).Return(nil, errors.New("test_error")).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			request, _ := http.NewRequest(http.MethodGet, "", nil)
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", test.token))
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
