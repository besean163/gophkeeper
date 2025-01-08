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

func TestUpdateNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := repositorymock.NewMockRepository(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	service := database.NewService(repository, encrypter, defaultlogger.NewDefaultLogger(), timecontroller, uuidController)
	user := models.User{ID: 1}
	item_1 := models.Note{
		Name:    "name_1",
		Content: "text_1",
	}
	item_2 := models.Note{
		UUID:    "uuid_2",
		Name:    "name_2",
		Content: "text_2",
	}

	tests := []struct {
		name      string
		item      models.Note
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			item: item_1,
			mockSetup: func() {
				uuidController.EXPECT().GetUUID().Return("new_uuid").Times(1)
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveNote(models.Note{
					UUID:      "new_uuid",
					UserID:    user.ID,
					Name:      item_1.Name,
					Content:   item_1.Content,
					UpdatedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "success_with_uuid",
			item: item_2,
			mockSetup: func() {
				timecontroller.EXPECT().Now().Return(int64(1)).Times(1)
				repository.EXPECT().SaveNote(models.Note{
					UUID:      item_2.UUID,
					UserID:    user.ID,
					Name:      item_2.Name,
					Content:   item_2.Content,
					UpdatedAt: 1,
				}).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.UpdateNote(user, test.item)
			assert.Equal(t, test.result, err)
		})
	}

}
