package bucket

import (
	"errors"
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	options := bucket.ServiceOptions{
		Repository:     repository,
		TimeController: timecontroller,
		UUIDController: uuidController,
	}
	service := bucket.NewService(options)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		mockSetup func()
		result    struct {
			items []*models.Card
			err   error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{
					{UUID: "uuid_1"},
					{UUID: "uuid_2"},
				}, nil).Times(1)
			},
			result: struct {
				items []*models.Card
				err   error
			}{
				items: []*models.Card{
					{UUID: "uuid_1"},
					{UUID: "uuid_2"},
				},
				err: nil,
			},
		},
		{
			name: "fail",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, errors.New("test_error")).Times(1)
			},
			result: struct {
				items []*models.Card
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
			items, err := service.GetCards(user)
			assert.Equal(t, test.result.items, items)
			assert.Equal(t, test.result.err, err)
		})
	}

}
