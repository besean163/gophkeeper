package api

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockDataService(ctrl)
	apiClient := mock.NewMockApiClient(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	syncController := mock.NewMockSyncer(ctrl)
	changeDetector := mock.NewMockChangeDetector(ctrl)

	service := api.NewService(storeService, apiClient, encrypter, timeController, nil, syncController, changeDetector)
	user := models.User{}

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncAccounts(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetAccounts().Return(&entities.GetAccountsOutput{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Account{{ID: 1}},
					[]models.Account{},
					[]models.Account{},
				).Times(1)
				storeService.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncAccounts(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetAccounts().Return(&entities.GetAccountsOutput{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Account{},
					[]models.Account{{ID: 1}},
					[]models.Account{},
				).Times(1)
				storeService.EXPECT().UpdateAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "delete",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncAccounts(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetAccounts().Return(&entities.GetAccountsOutput{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Account{},
					[]models.Account{},
					[]models.Account{{ID: 1}},
				).Times(1)
				storeService.EXPECT().DeleteAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncAccounts(user)
			assert.Equal(t, test.result, err)
		})
	}

}
