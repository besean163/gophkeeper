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
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAccountsSyncRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	user := models.User{Login: "user", Password: "password"}
	bucketService := mock.NewMockBucketService(ctrl)

	debs := dependencies.Dependencies{
		Logger:        defaultlogger.NewDefaultLogger(),
		BucketService: bucketService,
	}

	handler := route.AccountsSyncRoute(debs)

	tests := []struct {
		name         string
		exceptCode   int
		requestBody  string
		responseBody string
		user         *models.User
		mockSetup    func()
	}{
		{
			name:        "success",
			requestBody: `{"accounts":[{"uuid":"uuid_1","name":"name_1","login":"login_1","password":"password_1","created_at":1,"updated_at":1,"deleted_at":1,"synced_at":1}]}`,
			exceptCode:  http.StatusOK,
			user:        &user,
			mockSetup: func() {
				bucketService.EXPECT().SyncAccounts(gomock.Any(), gomock.Any(), []clientmodels.Account{
					{
						UUID:      "uuid_1",
						Name:      "name_1",
						Login:     "login_1",
						Password:  "password_1",
						CreatedAt: 1,
						UpdatedAt: 1,
						DeletedAt: 1,
						SyncedAt:  1,
					},
				}).Return(nil).Times(1)
			},
		},
		{
			name:         "not_auth",
			responseBody: `{"error":{"code":401,"description":"not authorized"}}`,
			exceptCode:   http.StatusUnauthorized,
			mockSetup:    func() {},
		},
		{
			name:         "unknown_error",
			requestBody:  `{"accounts":[{"uuid":"uuid_1","name":"name_1","login":"login_1","password":"password_1","created_at":1,"updated_at":1,"deleted_at":1,"synced_at":1}]}`,
			responseBody: `{"error":{"code":500,"description":"unknown internal error"}}`,
			exceptCode:   http.StatusInternalServerError,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().SyncAccounts(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test_error")).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			requestBody := strings.NewReader(test.requestBody)
			request, _ := http.NewRequest(http.MethodGet, "", requestBody)
			if test.user != nil {
				ctx := context.WithValue(context.Background(), entities.RequestUserKey("user"), test.user)
				request, _ = http.NewRequestWithContext(ctx, http.MethodGet, "", requestBody)
			}
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, request)

			response := rr.Result()
			defer response.Body.Close()

			// проверяем код ответа
			assert.Equal(t, test.exceptCode, response.StatusCode)

			// проверяем тело ответа
			responseBody, _ := io.ReadAll(response.Body)
			assert.Equal(t, test.responseBody, string(responseBody))
		})
	}
}
