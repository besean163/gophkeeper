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
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/models"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCardsSyncRoute(t *testing.T) {
	ctrl := gomock.NewController(t)

	user := models.User{Login: "user", Password: "password"}
	bucketService := mock.NewMockBucketService(ctrl)

	debs := dependencies.Dependencies{
		Logger:        defaultlogger.NewDefaultLogger(),
		BucketService: bucketService,
	}

	handler := route.CardsSyncRoute(debs)

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
			requestBody: `{"cards":[{"uuid":"uuid_1","name":"name_1","number":1111,"exp":"11|11","cvv":111,"created_at":1,"updated_at":1,"deleted_at":1,"synced_at":1}]}`,
			exceptCode:  http.StatusOK,
			user:        &user,
			mockSetup: func() {
				bucketService.EXPECT().SyncCards(gomock.Any(), gomock.Any(), []models.ExternalCard{
					{
						UUID:      "uuid_1",
						Name:      "name_1",
						Number:    1111,
						Exp:       "11|11",
						CVV:       111,
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
			requestBody:  `{"cards":[{"uuid":"uuid_1","name":"name_1","number":1111,"exp":"11|11","cvv":111,"created_at":1,"updated_at":1,"deleted_at":1,"synced_at":1}]}`,
			responseBody: `{"error":{"code":500,"description":"unknown internal error"}}`,
			exceptCode:   http.StatusInternalServerError,
			user:         &user,
			mockSetup: func() {
				bucketService.EXPECT().SyncCards(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test_error")).Times(1)
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
