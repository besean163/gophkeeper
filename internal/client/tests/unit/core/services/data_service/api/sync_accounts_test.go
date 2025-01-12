package api

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
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

	options := api.ServiceOptions{
		DataService:    storeService,
		ApiClient:      apiClient,
		Encrypter:      encrypter,
		TimeController: timeController,
		Syncer:         syncController,
		ChangeDetector: changeDetector,
	}
	service := api.NewService(options)

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
				apiClient.EXPECT().GetAccounts().Return(&output.GetAccounts{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any()).Return(
					changes.AccountChanges{
						Created: []models.Account{{UUID: "uuid"}},
						Updated: []models.Account{},
						Deleted: []models.Account{},
					},
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
				apiClient.EXPECT().GetAccounts().Return(&output.GetAccounts{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any()).Return(
					changes.AccountChanges{
						Created: []models.Account{},
						Updated: []models.Account{{UUID: "uuid"}},
						Deleted: []models.Account{},
					},
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
				apiClient.EXPECT().GetAccounts().Return(&output.GetAccounts{}, nil).Times(1)
				storeService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(2)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any()).Return(
					changes.AccountChanges{
						Created: []models.Account{},
						Updated: []models.Account{},
						Deleted: []models.Account{{UUID: "uuid"}},
					},
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
