package api

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockDataService(ctrl)
	apiClient := mock.NewMockApiClient(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	syncController := mock.NewMockSyncer(ctrl)

	service := api.NewService(storeService, apiClient, encrypter, timeController, nil, syncController, nil)

	tests := []struct {
		name      string
		user      models.User
		item      models.Account
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			user: models.User{},
			item: models.Account{},
			mockSetup: func() {
				storeService.EXPECT().DeleteAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				syncController.EXPECT().SyncAccounts(gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.DeleteAccount(test.user, test.item, false)
			assert.Equal(t, test.result, err)
		})
	}

}
