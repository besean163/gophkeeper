package auth

import (
	"errors"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/services/auth"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)

	service := auth.NewService(repository, nil, nil, nil, nil)

	tests := []struct {
		name      string
		id        int
		mockSetup func()
		result    struct {
			user *models.User
			err  error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetUser(gomock.Any()).Return(&models.User{}, nil).Times(1)
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
			name: "fail",
			mockSetup: func() {
				repository.EXPECT().GetUser(gomock.Any()).Return(nil, errors.New("test_error")).Times(1)
			},
			result: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("test_error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			user, err := service.GetUser(1)
			assert.Equal(t, test.result.user, user)
			assert.Equal(t, test.result.err, err)
		})
	}
}
