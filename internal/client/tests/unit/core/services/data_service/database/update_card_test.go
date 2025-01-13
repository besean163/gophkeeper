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

func TestUpdateCard(t *testing.T) {
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
	item_1 := models.Card{
		Name:   "name_1",
		Number: 1111,
		Exp:    "11|11",
		CVV:    111,
	}
	item_2 := models.Card{
		UUID:   "uuid_2",
		Name:   "name_2",
		Number: 2222,
		Exp:    "22|22",
		CVV:    222,
	}

	tests := []struct {
		name      string
		item      models.Card
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: item_1,
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveCard(models.Card{
					UserID:    user.ID,
					Name:      item_1.Name,
					Number:    item_1.Number,
					Exp:       item_1.Exp,
					CVV:       item_1.CVV,
					UpdatedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "success_with_uuid",
			item: item_2,
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveCard(models.Card{
					UUID:      item_2.UUID,
					UserID:    user.ID,
					Name:      item_2.Name,
					Number:    item_2.Number,
					Exp:       item_2.Exp,
					CVV:       item_2.CVV,
					UpdatedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.UpdateCard(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
