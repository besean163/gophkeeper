package models

import (
	"testing"

	coremodels "github.com/besean163/gophkeeper/internal/models/client"

	mock "github.com/besean163/gophkeeper/internal/client/tests/mocks"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/golang/mock/gomock"
)

func TestCardsListModel_init(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)

	model := models.NewCardListModel(core, defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestCardsListModel_update(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)

	model := models.NewCardListModel(core, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		msg       tea.Msg
		mockSetup func()
	}{
		{
			name: "enter",
			msg:  tea.KeyMsg{Type: tea.KeyEnter},
			mockSetup: func() {
				core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)
			},
		},
		{
			name:      "ctrl+b",
			msg:       tea.KeyMsg{Type: tea.KeyTab},
			mockSetup: func() {},
		},
		{
			name:      "ctrl+a",
			msg:       tea.KeyMsg{Type: tea.KeyShiftTab},
			mockSetup: func() {},
		},
		{
			name:      "ctrl+d",
			msg:       tea.KeyMsg{Type: tea.KeyUp},
			mockSetup: func() {},
		},
		{
			name: "delete",
			msg:  messages.CardDeleteMsg{},
			mockSetup: func() {
				core.EXPECT().DeleteCard(coremodels.Card{}).Return(nil).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			model.Update(test.msg)
		})
	}
}

func TestCardsListModel_view(t *testing.T) {
	ctrl := gomock.NewController(t)
	core := mock.NewMockCore(ctrl)
	core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)

	model := models.NewCardListModel(core, defaultlogger.NewDefaultLogger())

	model.View()
}
