package bucket

import (
	"errors"
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/services/bucket"
	bucketmock "github.com/besean163/gophkeeper/internal/server/tests/mocks"
	servicemock "github.com/besean163/gophkeeper/internal/server/tests/mocks/services/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := servicemock.NewMockRepository(ctrl)
	changeDetector := servicemock.NewMockChangeDetector(ctrl)
	timecontroller := utilmock.NewMockTimeController(ctrl)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	selfService := bucketmock.NewMockBucketService(ctrl)

	service := bucket.NewService(repository, timecontroller, uuidController, changeDetector)
	user := models.User{ID: 1}

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{{UUID: "uuid_1"}},
					[]*models.Account{},
					[]string{},
				).Times(1)
				selfService.EXPECT().CreateAccount(user, &models.Account{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "create_fail",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{{UUID: "uuid_1"}},
					[]*models.Account{},
					[]string{},
				).Times(1)
				selfService.EXPECT().CreateAccount(user, &models.Account{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "update",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{},
					[]*models.Account{{UUID: "uuid_1"}},
					[]string{},
				).Times(1)
				selfService.EXPECT().UpdateAccount(user, &models.Account{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "update_fail",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{},
					[]*models.Account{{UUID: "uuid_1"}},
					[]string{},
				).Times(1)
				selfService.EXPECT().UpdateAccount(user, &models.Account{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "delete",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{},
					[]*models.Account{},
					[]string{"uuid_1"},
				).Times(1)
				selfService.EXPECT().DeleteAccount(user, "uuid_1").Return(nil).Times(1)
			},
		},
		{
			name: "delete_fail",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{},
					[]*models.Account{},
					[]string{"uuid_1"},
				).Times(1)
				selfService.EXPECT().DeleteAccount(user, "uuid_1").Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "ignore",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, nil).Times(1)
				changeDetector.EXPECT().GetAccountChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]*models.Account{},
					[]*models.Account{},
					[]string{},
				).Times(1)
			},
		},
		{
			name: "get_fail",
			mockSetup: func() {
				repository.EXPECT().GetAccounts(user).Return([]*models.Account{}, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncAccounts(selfService, user, []models.ExternalAccount{})
			assert.Equal(t, test.result, err)
		})
	}
}
