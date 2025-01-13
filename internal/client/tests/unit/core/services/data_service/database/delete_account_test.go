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

func TestDeleteAccount(t *testing.T) {
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
	item_1 := models.Account{
		Name:     "name_1",
		Login:    "login_1",
		Password: "password_1",
	}

	tests := []struct {
		name      string
		soft      bool
		item      models.Account
		mockSetup func()
		result    error
	}{
		{
			name: "success_soft",
			soft: true,
			item: item_1,
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveAccount(models.Account{
					Name:      item_1.Name,
					Login:     item_1.Login,
					Password:  item_1.Password,
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
				repository.EXPECT().DeleteAccount("").Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.DeleteAccount(user, test.item, test.soft)
			assert.Equal(t, test.result, err)
		})
	}

}
