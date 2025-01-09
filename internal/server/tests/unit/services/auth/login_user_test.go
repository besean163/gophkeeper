package auth

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/services/auth"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/auth"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	tokener := utilmock.NewMockTokener(ctrl)
	repository := repositorymock.NewMockRepository(ctrl)

	options := auth.ServiceOptions{
		Repository: repository,
		Encrypter:  encrypter,
		Tokener:    tokener,
	}
	service := auth.NewService(options)

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
				repository.EXPECT().GetUserByLogin("login").Return(&models.User{}).Times(1)
				encrypter.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(true).Times(1)
				tokener.EXPECT().GetToken(gomock.Any()).Return("token", nil).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "token", err: nil},
		},
		{
			name:     "user_not_found",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin("login").Return(nil).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "", err: apierrors.ErrorUserNotExist},
		},
		{
			name:     "wrong_password",
			login:    "login",
			password: "password",
			mockSetup: func() {
				repository.EXPECT().GetUserByLogin("login").Return(&models.User{}).Times(1)
				encrypter.EXPECT().CheckPassword(gomock.Any(), gomock.Any()).Return(false).Times(1)
			},
			result: struct {
				token string
				err   error
			}{token: "", err: apierrors.ErrorValidatePasswordWrong},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			token, err := service.LoginUser(test.login, test.password)
			assert.Equal(t, test.result.token, token)
			assert.Equal(t, test.result.err, err)
		})
	}

}
