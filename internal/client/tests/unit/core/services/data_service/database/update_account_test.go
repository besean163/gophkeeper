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

func TestUpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := database.NewService(repository, encrypter, defaultlogger.NewDefaultLogger(), timecontroller, uuidController)
	user := models.User{ID: 1}
	item_1 := models.Account{
		Name:     "name_1",
		Login:    "login_1",
		Password: "password_1",
	}
	item_2 := models.Account{
		UUID:     "uuid_2",
		Name:     "name_2",
		Login:    "login_2",
		Password: "password_2",
	}

	tests := []struct {
		name      string
		item      models.Account
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: item_1,
			mockSetup: func() {
				uuidController.EXPECT().GetUUID().Return("new_uuid").Times(1)
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveAccount(models.Account{
					UUID:      "new_uuid",
					UserID:    user.ID,
					Name:      item_1.Name,
					Login:     item_1.Login,
					Password:  item_1.Password,
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
				repository.EXPECT().SaveAccount(models.Account{
					UUID:      item_2.UUID,
					UserID:    user.ID,
					Name:      item_2.Name,
					Login:     item_2.Login,
					Password:  item_2.Password,
					UpdatedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.UpdateAccount(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
