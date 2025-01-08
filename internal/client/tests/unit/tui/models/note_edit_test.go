package models

import (
	"testing"

	coremodels "github.com/besean163/gophkeeper/internal/client/core/models"
	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/golang/mock/gomock"
)

func TestNoteEditModel_init(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewNoteEditModel(core, coremodels.Note{}, defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestNoteEditModel_update(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewNoteEditModel(core, coremodels.Note{}, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		msg       tea.Msg
		mockSetup func()
	}{
		{
			name:      "tab",
			msg:       tea.KeyMsg{Type: tea.KeyTab},
			mockSetup: func() {},
		},
		{
			name:      "shift+tab",
			msg:       tea.KeyMsg{Type: tea.KeyShiftTab},
			mockSetup: func() {},
		},
		{
			name:      "enter",
			msg:       tea.KeyMsg{Type: tea.KeyEnter},
			mockSetup: func() {},
		},
		{
			name: "submit",
			msg:  messages.ButtonSubmitMsg{},
			mockSetup: func() {
				core.EXPECT().SaveNote(gomock.Any()).Return(nil).Times(1)
			},
		},
		{
			name:      "back",
			msg:       messages.ButtonBackMsg{},
			mockSetup: func() {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			model.Update(test.msg)
		})
	}
}

func TestNoteEditModel_view(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewNoteEditModel(core, coremodels.Note{}, defaultlogger.NewDefaultLogger())

	model.View()
}
