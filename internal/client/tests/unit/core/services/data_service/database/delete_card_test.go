package bucket

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/database"
	repositorymock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := database.NewService(repository, encrypter, defaultlogger.NewDefaultLogger(), timecontroller, uuidController)
	user := models.User{ID: 1}
	item_1 := models.Card{
		Name:   "name_1",
		Number: 1111,
		Exp:    "11|11",
		CVV:    111,
	}

	tests := []struct {
		name      string
		soft      bool
		item      models.Card
		mockSetup func()
		result    error
	}{
		{
			name: "success_soft",
			soft: true,
			item: item_1,
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveCard(models.Card{
					Name:      item_1.Name,
					Number:    item_1.Number,
					Exp:       item_1.Exp,
					CVV:       item_1.CVV,
					DeletedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "success",
			soft: false,
			item: item_1,
			mockSetup: func() {
				repository.EXPECT().DeleteCard("").Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.DeleteCard(user, test.item, test.soft)
			assert.Equal(t, test.result, err)
		})
	}

}
