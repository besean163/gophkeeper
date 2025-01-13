package bucket

import (
	"errors"
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/database"
	repositorymock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
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

	tests := []struct {
		name      string
		mockSetup func()
		result    struct {
			user *models.User
			err  error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(&models.User{}).Times(1)
				encrypter.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: &models.User{},
				err:  nil,
			},
		},
		{
			name: "not_found_user",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(nil).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("user not exist"),
			},
		},
		{
			name: "wrong_password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(&models.User{}).Times(1)
				encrypter.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("wrong password"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			user, err := service.LoginUser("test_login", "test_password")
			assert.Equal(t, test.result.user, user)
			assert.Equal(t, test.result.err, err)
		})
	}

}
