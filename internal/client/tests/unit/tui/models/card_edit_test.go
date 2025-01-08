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

func TestCardEditModel_init(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewCardEditModel(core, coremodels.Card{}, defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestCardEditModel_update(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewCardEditModel(core, coremodels.Card{}, defaultlogger.NewDefaultLogger())

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
			name:      "up",
			msg:       tea.KeyMsg{Type: tea.KeyUp},
			mockSetup: func() {},
		},
		{
			name:      "down",
			msg:       tea.KeyMsg{Type: tea.KeyDown},
			mockSetup: func() {},
		},
		{
			name: "submit",
			msg:  messages.ButtonSubmitMsg{},
			mockSetup: func() {
				core.EXPECT().SaveCard(gomock.Any()).Return(nil).Times(1)
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

func TestCardEditModel_view(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	model := models.NewCardEditModel(core, coremodels.Card{}, defaultlogger.NewDefaultLogger())

	model.View()
}
