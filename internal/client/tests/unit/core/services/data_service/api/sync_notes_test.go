package api

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"

	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeService := mock.NewMockDataService(ctrl)
	apiClient := mock.NewMockApiClient(ctrl)
	encrypter := utilmock.NewMockEncrypter(ctrl)
	timeController := utilmock.NewMockTimeController(ctrl)
	syncController := mock.NewMockSyncer(ctrl)
	changeDetector := mock.NewMockChangeDetector(ctrl)

	options := api.ServiceOptions{
		DataService:    storeService,
		ApiClient:      apiClient,
		Encrypter:      encrypter,
		TimeController: timeController,
		Syncer:         syncController,
		ChangeDetector: changeDetector,
	}
	service := api.NewService(options)

	user := models.User{}

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncNotes(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetNotes().Return(&output.GetNotes{}, nil).Times(1)
				storeService.EXPECT().GetNotes(gomock.Any()).Return([]models.Note{}, nil).Times(2)
				changeDetector.EXPECT().GetNoteChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Note{{UUID: "uuid"}},
					[]models.Note{},
					[]models.Note{},
				).Times(1)
				storeService.EXPECT().CreateNote(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncNotes(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetNotes().Return(&output.GetNotes{}, nil).Times(1)
				storeService.EXPECT().GetNotes(gomock.Any()).Return([]models.Note{}, nil).Times(2)
				changeDetector.EXPECT().GetNoteChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Note{},
					[]models.Note{{UUID: "uuid"}},
					[]models.Note{},
				).Times(1)
				storeService.EXPECT().UpdateNote(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "delete",
			mockSetup: func() {
				apiClient.EXPECT().HasConnection().Return(true).Times(2)
				apiClient.EXPECT().SetToken(gomock.Any()).Times(2)
				apiClient.EXPECT().SyncNotes(gomock.Any()).Return(nil).Times(1)
				apiClient.EXPECT().GetNotes().Return(&output.GetNotes{}, nil).Times(1)
				storeService.EXPECT().GetNotes(gomock.Any()).Return([]models.Note{}, nil).Times(2)
				changeDetector.EXPECT().GetNoteChanges(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					[]models.Note{},
					[]models.Note{},
					[]models.Note{{UUID: "uuid"}},
				).Times(1)
				storeService.EXPECT().DeleteNote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncNotes(user)
			assert.Equal(t, test.result, err)
		})
	}

}
