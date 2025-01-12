package bucket

import (
	"errors"
	"testing"

	clientmodels "github.com/besean163/gophkeeper/internal/models/client"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
	bucketmock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	servicemock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := servicemock.NewMockRepository(ctrl)
	changeDetector := servicemock.NewMockChangeDetector(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	selfService := bucketmock.NewMockBucketService(ctrl)

	options := bucket.ServiceOptions{
		Repository:     repository,
		TimeController: timecontroller,
		UUIDController: uuidController,
		ChangeDetector: changeDetector,
	}
	service := bucket.NewService(options)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{{UUID: "uuid_1"}},
						Updated: []*models.Card{},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().CreateCard(user, &models.Card{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "create_fail",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{{UUID: "uuid_1"}},
						Updated: []*models.Card{},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().CreateCard(user, &models.Card{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "update",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{},
						Updated: []*models.Card{{UUID: "uuid_1"}},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().UpdateCard(user, &models.Card{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "update_fail",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{},
						Updated: []*models.Card{{UUID: "uuid_1"}},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().UpdateCard(user, &models.Card{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "delete",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{},
						Updated: []*models.Card{},
						Deleted: []string{"uuid_1"},
					},
				).Times(1)
				selfService.EXPECT().DeleteCard(user, "uuid_1").Return(nil).Times(1)
			},
		},
		{
			name: "delete_fail",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{},
						Updated: []*models.Card{},
						Deleted: []string{"uuid_1"},
					},
				).Times(1)
				selfService.EXPECT().DeleteCard(user, "uuid_1").Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "ignore",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, nil).Times(1)
				changeDetector.EXPECT().GetCardsChanges(gomock.Any(), gomock.Any()).Return(
					changes.CardChanges{
						Created: []*models.Card{},
						Updated: []*models.Card{},
						Deleted: []string{},
					},
				).Times(1)
			},
		},
		{
			name: "get_fail",
			mockSetup: func() {
				repository.EXPECT().GetCards(user).Return([]*models.Card{}, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncCards(selfService, user, []clientmodels.Card{})
			assert.Equal(t, test.result, err)
		})
	}
}
