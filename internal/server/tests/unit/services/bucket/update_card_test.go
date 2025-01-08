package bucket

import (
	"errors"
	"testing"

	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := bucket.NewService(repository, timecontroller, uuidController, nil)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		item      *models.Card
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: &models.Card{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetCard("uuid_1").Return(&models.Card{UserID: 1, UUID: "uuid_1", Name: "name_1"}, nil).Times(1)
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveCard(&models.Card{UserID: 1, UUID: "uuid_1", Name: "new_name_1", UpdatedAt: 1}).Return(nil).Times(1)
			},
		},
		{
			name: "not_found",
			item: &models.Card{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetCard("uuid_1").Return(nil, nil).Times(1)
			},
			result: apierrors.ErrorNotFoundByUUID,
		},
		{
			name: "fail",
			item: &models.Card{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetCard("uuid_1").Return(nil, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.UpdateCard(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
