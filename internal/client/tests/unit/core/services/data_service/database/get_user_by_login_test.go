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

func TestGetUserByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := database.NewService(repository, encrypter, defaultlogger.NewDefaultLogger(), timecontroller, uuidController)

	tests := []struct {
		name      string
		mockSetup func()
		result    *models.User
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(&models.User{}).Times(1)
			},
			result: &models.User{},
		},
		{
			name: "fail",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			user := service.GetUserByLogin("test_login")
			assert.Equal(t, test.result, user)
		})
	}

}
