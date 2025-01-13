package api

import (
	"errors"
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockDataService(ctrl)
	apiClient := mock.NewMockApiClient(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	syncController := mock.NewMockSyncer(ctrl)

	options := api.ServiceOptions{
		DataService:    storeService,
		ApiClient:      apiClient,
		Encrypter:      encrypter,
		TimeController: timeController,
		Syncer:         syncController,
	}
	service := api.NewService(options)

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
				apiClient.EXPECT().Register(gomock.Any()).Return(output.Token{Token: "token"}, nil).Times(1)
				encrypter.EXPECT().Encrypt(gomock.Any()).Return("encrypted_password", nil)
				storeService.EXPECT().SaveUser(models.User{
					Login:    "login",
					Password: "encrypted_password",
					Token:    "token",
				}).Return(nil).Times(1)
				syncController.EXPECT().Sync(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: &models.User{
					Login:    "login",
					Password: "encrypted_password",
					Token:    "token",
				},
				err: nil,
			},
		},
		{
			name:     "no_connection",
			login:    "login",
			password: "password",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(false).Times(1)
				// storeService.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("can't register user offline")).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("can't register user offline"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			user, err := service.RegisterUser(test.login, test.password)
			assert.Equal(t, test.result.user, user)
			assert.Equal(t, test.result.err, err)
		})
	}

}
