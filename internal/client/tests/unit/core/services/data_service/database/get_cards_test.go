package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/database"
	repositorymock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	options := database.ServiceOptions{
		Repository:     repository,
		Encrypter:      encrypter,
		Logger:         defaultlogger.NewDefaultLogger(),
		TimeController: timecontroller,
		UUIDController: uuidController,
	}

	service := database.NewService(options)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		mockSetup func()
		result    struct {
			items []models.Card
			err   error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetCards(gomock.Any()).Return([]models.Card{}, nil).Times(1)
			},
			result: struct {
				items []models.Card
				err   error
			}{
				items: []models.Card{},
				err:   nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			items, err := service.GetCards(user)
			assert.Equal(t, test.result.items, items)
			assert.Equal(t, test.result.err, err)
		})
	}

}
