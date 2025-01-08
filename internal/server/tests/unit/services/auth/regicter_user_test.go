package auth

import (
	"errors"
	"testing"

	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/services/auth"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/auth"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	tokener := utilmock.NewMockTokener(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := auth.NewService(repository, encrypter, tokener, timeController, uuidController)

	tests := []struct {
		name      string
		login     string
		password  string
		mockSetup func()
		result    struct {
			token string
			err   error
		}
	}{
		{
			name:     "success",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin("login").Return(nil).Times(1)
				encrypter.EXPECT().Encrypt(gomock.Any()).Return("encrypted_password", nil).Times(1)
				uuidController.EXPECT().GetUUID().Return("new_uuid").Times(1)
				timeController.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveUser(&models.User{
					UUID:      "new_uuid",
					Login:     "login",
					Password:  "encrypted_password",
					CreatedAt: 1,
				}).Return(nil)
				tokener.EXPECT().GetToken(gomock.Any()).Return("token", nil).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "token", err: nil},
		},
		{
			name:     "user_exist",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin("login").Return(&models.User{}).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "", err: apierrors.ErrorUserExist},
		},
		{
			name:     "password_encrypt_error",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin("login").Return(nil).Times(1)
				encrypter.EXPECT().Encrypt(gomock.Any()).Return("", errors.New("test_error")).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "", err: errors.New("test_error")},
		},
		{
			name:     "save_user_error",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin(gomock.Any()).Return(nil).Times(1)
				encrypter.EXPECT().Encrypt(gomock.Any()).Return("", nil).Times(1)
				uuidController.EXPECT().GetUUID().Return("").Times(1)
				timeController.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveUser(gomock.Any()).Return(errors.New("test_error"))
			},
			result: struct {
				token string
				err   error
			}{token: "", err: errors.New("test_error")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			token, err := service.RegisterUser(test.login, test.password)
			assert.Equal(t, test.result.token, token)
			assert.Equal(t, test.result.err, err)
		})
	}

}
