package api

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCard(t *testing.T) {
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
		user      models.User
		item      models.Card
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			user: models.User{},
			item: models.Card{},
			mockSetup: func() {
				storeService.EXPECT().DeleteCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				syncController.EXPECT().Sync(gomock.Any(),gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.DeleteCard(test.user, test.item, false)
			assert.Equal(t, test.result, err)
		})
	}

}
