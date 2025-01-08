package api

import (
	"errors"
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockDataService(ctrl)
	apiClient := mock.NewMockApiClient(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	syncController := mock.NewMockSyncer(ctrl)

	service := api.NewService(storeService, apiClient, encrypter, timeController, nil, syncController, nil)

	tests := []struct {
		name      string
		login     string
		password  string
		mockSetup func()
		result    struct {
			user *models.User
			err  error
		}
	}{
		{
			name:     "success",
			login:    "login",
			password: "password",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(1)
				apiClient.EXPECT().Login(gomock.Any()).Return(entities.TokenOutput{Token: "token"}, nil).Times(1)
				encrypter.EXPECT().Encrypt(gomock.Any()).Return("encrypted_password", nil)
				storeService.EXPECT().SaveUser(models.User{
					Login:    "login",
					Password: "encrypted_password",
					Token:    "token",
				}).Return(nil).Times(1)
				storeService.EXPECT().GetUserByLogin(gomock.Any()).Return(&models.User{}).Times(1)
				syncController.EXPECT().SyncAll(gomock.Any()).Return(nil).Times(1)
				storeService.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(&models.User{}, nil).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: &models.User{},
				err:  nil,
			},
		},
		{
			name:     "api_login_error",
			login:    "login",
			password: "password",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(1)
				apiClient.EXPECT().Login(gomock.Any()).Return(entities.TokenOutput{}, errors.New("test_error")).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("test_error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			user, err := service.LoginUser(test.login, test.password)
			assert.Equal(t, test.result.user, user)
			assert.Equal(t, test.result.err, err)
		})
	}

}
