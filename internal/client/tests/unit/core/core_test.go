package core

import (
	"errors"
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core"
	models "github.com/besean163/gophkeeper/internal/models/client"

	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(&models.User{}, nil).Times(1)
			},
			result: nil,
		},
		{
			name: "fail",
			mockSetup: func() {
				dataService.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := c.Login("", "")
			assert.Equal(t, test.result, err)
		})
	}
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(&models.User{}, nil).Times(1)
			},
			result: nil,
		},
		{
			name: "fail",
			mockSetup: func() {
				dataService.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			err := c.Register("", "")
			assert.Equal(t, test.result, err)
		})
	}
}

func TestGetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().GetAccounts(gomock.Any()).Return([]models.Account{}, nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			_, err := c.GetAccounts()
			assert.Equal(t, test.result, err)
		})
	}
}

func TestSaveAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		item      models.Account
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			item: models.Account{},
			mockSetup: func() {
				dataService.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			item: models.Account{UUID: "uuid"},
			mockSetup: func() {
				dataService.EXPECT().UpdateAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.SaveAccount(test.item)
			assert.Equal(t, test.result, err)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().DeleteAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "fail",
			mockSetup: func() {
				dataService.EXPECT().DeleteAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.DeleteAccount(models.Account{})
			assert.Equal(t, test.result, err)
		})
	}
}

func TestGetNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().GetNotes(gomock.Any()).Return([]models.Note{}, nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			_, err := c.GetNotes()
			assert.Equal(t, test.result, err)
		})
	}
}

func TestSaveNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		item      models.Note
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			item: models.Note{},
			mockSetup: func() {
				dataService.EXPECT().CreateNote(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			item: models.Note{UUID: "uuid"},
			mockSetup: func() {
				dataService.EXPECT().UpdateNote(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.SaveNote(test.item)
			assert.Equal(t, test.result, err)
		})
	}
}

func TestDeleteNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().DeleteNote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "fail",
			mockSetup: func() {
				dataService.EXPECT().DeleteNote(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.DeleteNote(models.Note{})
			assert.Equal(t, test.result, err)
		})
	}
}

func TestGetCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().GetCards(gomock.Any()).Return([]models.Card{}, nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			_, err := c.GetCards()
			assert.Equal(t, test.result, err)
		})
	}
}

func TestSaveCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		item      models.Card
		mockSetup func()
		result    error
	}{
		{
			name: "create",
			item: models.Card{},
			mockSetup: func() {
				dataService.EXPECT().CreateCard(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "update",
			item: models.Card{UUID: "uuid"},
			mockSetup: func() {
				dataService.EXPECT().UpdateCard(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.SaveCard(test.item)
			assert.Equal(t, test.result, err)
		})
	}
}

func TestDeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	dataService := mock.NewMockDataService(ctrl)

	c := core.NewCore(dataService, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				dataService.EXPECT().DeleteCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			result: nil,
		},
		{
			name: "fail",
			mockSetup: func() {
				dataService.EXPECT().DeleteCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test_error")).Times(1)
			},
			result: errors.New("test_error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.User = &models.User{}
			err := c.DeleteCard(models.Card{})
			assert.Equal(t, test.result, err)
		})
	}
}
