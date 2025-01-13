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

func TestSyncNotes(t *testing.T) {
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
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{{UUID: "uuid_1"}},
						Updated: []*models.Note{},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().CreateNote(user, &models.Note{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "create_fail",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{{UUID: "uuid_1"}},
						Updated: []*models.Note{},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().CreateNote(user, &models.Note{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "update",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{},
						Updated: []*models.Note{{UUID: "uuid_1"}},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().UpdateNote(user, &models.Note{UUID: "uuid_1"}).Return(nil).Times(1)
			},
		},
		{
			name: "update_fail",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{},
						Updated: []*models.Note{{UUID: "uuid_1"}},
						Deleted: []string{},
					},
				).Times(1)
				selfService.EXPECT().UpdateNote(user, &models.Note{UUID: "uuid_1"}).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "delete",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{},
						Updated: []*models.Note{},
						Deleted: []string{"uuid_1"},
					},
				).Times(1)
				selfService.EXPECT().DeleteNote(user, "uuid_1").Return(nil).Times(1)
			},
		},
		{
			name: "delete_fail",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{},
						Updated: []*models.Note{},
						Deleted: []string{"uuid_1"},
					},
				).Times(1)
				selfService.EXPECT().DeleteNote(user, "uuid_1").Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
		{
			name: "ignore",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, nil).Times(1)
				changeDetector.EXPECT().GetNotesChanges(gomock.Any(), gomock.Any()).Return(
					changes.NoteChanges{
						Created: []*models.Note{},
						Updated: []*models.Note{},
						Deleted: []string{},
					},
				).Times(1)
			},
		},
		{
			name: "get_fail",
			mockSetup: func() {
				repository.EXPECT().GetNotes(user).Return([]*models.Note{}, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := service.SyncNotes(selfService, user, []clientmodels.Note{})
			assert.Equal(t, test.result, err)
		})
	}
}
