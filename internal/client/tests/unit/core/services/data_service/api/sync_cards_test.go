package api

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncCards(t *testing.T) {
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
				apiClient.EXPECT().SyncCards(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetCards().Return(&output.GetCards{}, nil).Times(1)
				storeService.EXPECT().GetCards(gomock.Any()).Return([]models.Card{}, nil).Times(2)
				changeDetector.EXPECT().GetCardChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Card{{UUID: "uuid"}},
					[]models.Card{},
					[]models.Card{},
				).Times(1)
				storeService.EXPECT().CreateCard(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncCards(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetCards().Return(&output.GetCards{}, nil).Times(1)
				storeService.EXPECT().GetCards(gomock.Any()).Return([]models.Card{}, nil).Times(2)
				changeDetector.EXPECT().GetCardChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Card{},
					[]models.Card{{UUID: "uuid"}},
					[]models.Card{},
				).Times(1)
				storeService.EXPECT().UpdateCard(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "delete",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncCards(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetCards().Return(&output.GetCards{}, nil).Times(1)
				storeService.EXPECT().GetCards(gomock.Any()).Return([]models.Card{}, nil).Times(2)
				changeDetector.EXPECT().GetCardChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Card{},
					[]models.Card{},
					[]models.Card{{UUID: "uuid"}},
				).Times(1)
				storeService.EXPECT().DeleteCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncCards(user)
			assert.Equal(t, test.result, err)
		})
	}

}
