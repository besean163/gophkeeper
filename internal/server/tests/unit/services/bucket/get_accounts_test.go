package bucket

import (
	"errors"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := bucket.NewService(repository, timecontroller, uuidController, nil)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		mockSetup func()
		result    struct {
			items []*models.Account
			err   error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{
					{ID: 1},
					{ID: 2},
				}, nil).Times(1)
			},
			result: struct {
				items []*models.Account
				err   error
			}{
				items: []*models.Account{
					{ID: 1},
					{ID: 2},
				},
				err: nil,
			},
		},
		{
			name: "fail",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, errors.New("test_error")).Times(1)
			},
			result: struct {
				items []*models.Account
				err   error
			}{
				items: nil,
				err:   errors.New("test_error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			items, err := service.GetAccounts(user)
			assert.Equal(t, test.result.items, items)
			assert.Equal(t, test.result.err, err)
		})
	}

}
