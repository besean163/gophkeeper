package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	repositorymock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
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
				timecontroller.EXPECT().Now().Return(int64(2)).Times(1)
				repository.EXPECT().SaveNote(&models.Note{UserID: user.ID, UUID: "some_uuid", CreatedAt: 2}).Return(nil).Times(1)
			},
		},
		{
			name: "without_uuid",
			item: &models.Note{},
			mockSetup: func() {
				uuidController.EXPECT().GetUUID().Return("new_uuid").Times(1)
				timecontroller.EXPECT().Now().Return(int64(2)).Times(1)
				repository.EXPECT().SaveNote(&models.Note{UserID: user.ID, UUID: "new_uuid", CreatedAt: 2}).Return(nil).Times(1)
			},
		},
		{
			name: "without_user_id",
			item: &models.Note{UUID: "some_uuid"},
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(2)).Times(1)
				repository.EXPECT().SaveNote(&models.Note{UserID: user.ID, UUID: "some_uuid", CreatedAt: 2}).Return(nil).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.CreateNote(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
