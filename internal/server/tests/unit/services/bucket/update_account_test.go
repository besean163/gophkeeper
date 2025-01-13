package bucket

import (
	"errors"
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccount(t *testing.T) {
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
		item      *models.Account
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: &models.Account{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetAccount("uuid_1").Return(&models.Account{UserID: 1, UUID: "uuid_1", Name: "name_1"}, nil).Times(1)
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveAccount(&models.Account{UserID: 1, UUID: "uuid_1", Name: "new_name_1", UpdatedAt: 1}).Return(nil).Times(1)
			},
		},
		{
			name: "not_found",
			item: &models.Account{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetAccount("uuid_1").Return(nil, nil).Times(1)
			},
			result: apierrors.ErrorNotFoundByUUID,
		},
		{
			name: "fail",
			item: &models.Account{UserID: 1, UUID: "uuid_1", Name: "new_name_1"},
			mockSetup: func() {
				repository.EXPECT().GetAccount("uuid_1").Return(nil, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.UpdateAccount(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
