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

func TestDeleteNote(t *testing.T) {
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
		item      *models.Note
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: &models.Note{UserID: 1, UUID: "some_uuid"},
			mockSetup: func() {
				repository.EXPECT().DeleteNote("some_uuid").Return(nil).Times(1)
			},
		},
		{
			name: "fail",
			item: &models.Note{UserID: 1, UUID: "some_uuid"},
			mockSetup: func() {
				repository.EXPECT().DeleteNote("some_uuid").Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.DeleteNote(user, test.item.UUID)
			assert.Equal(t, test.result, err)
		})
	}
}
