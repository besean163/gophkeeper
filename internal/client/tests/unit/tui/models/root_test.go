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

func TestRoot_init(t *testing.T) {
	model := models.NewRootModel(nil, defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestRoot_update(t *testing.T) {
	ctrl := gomock.NewController(t)

	core := mock.NewMockCore(ctrl)
	model := models.NewRootModel(core, defaultlogger.NewDefaultLogger())
	tests := []struct {
		name      string
		msg       tea.Msg
		mockSetup func()
	}{
		{
			name:      "choose login",
			msg:       messages.SignLoginMsg{},
			mockSetup: func() {},
		},
		{
			name:      "choose registration",
			msg:       messages.SignRegistrationMsg{},
			mockSetup: func() {},
		},
		{
			name:      "return to sign",
			msg:       messages.SignBackMsg{},
			mockSetup: func() {},
		},
		{
			name:      "success login",
			msg:       messages.LoginSuccessMsg{},
			mockSetup: func() {},
		},
		{
			name:      "success registration",
			msg:       messages.RegistrationSuccessMsg{},
			mockSetup: func() {},
		},
		{
			name:      "return to select section",
			msg:       messages.SectionBackMsg{},
			mockSetup: func() {},
		},
		{
			name: "to account list",
			msg:  messages.SelectAccountMsg{},
			mockSetup: func() {
				core.EXPECT().GetAccounts().Return([]coremodels.Account{}, nil).Times(1)
			},
		},
		{
			name: "return to account list",
			msg:  messages.AccountListBackMsg{},
			mockSetup: func() {
				core.EXPECT().GetAccounts().Return([]coremodels.Account{}, nil).Times(1)
			},
		},
		{
			name:      "to account edit",
			msg:       messages.AccountEditMsg{Account: coremodels.Account{}},
			mockSetup: func() {},
		},
		{
			name: "to note list",
			msg:  messages.SelectNoteMsg{},
			mockSetup: func() {
				core.EXPECT().GetNotes().Return([]coremodels.Note{}, nil).Times(1)
			},
		},
		{
			name: "return to note list",
			msg:  messages.NoteListBackMsg{},
			mockSetup: func() {
				core.EXPECT().GetNotes().Return([]coremodels.Note{}, nil).Times(1)
			},
		},
		{
			name:      "to note edit",
			msg:       messages.NoteEditMsg{Note: coremodels.Note{}},
			mockSetup: func() {},
		},
		{
			name: "to card list",
			msg:  messages.SelectCardMsg{},
			mockSetup: func() {
				core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)
			},
		},
		{
			name: "return to card list",
			msg:  messages.CardListBackMsg{},
			mockSetup: func() {
				core.EXPECT().GetCards().Return([]coremodels.Card{}, nil).Times(1)
			},
		},
		{
			name:      "to card edit",
			msg:       messages.CardEditMsg{Card: coremodels.Card{}},
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

func TestRoot_view(t *testing.T) {
	model := models.NewRootModel(nil, defaultlogger.NewDefaultLogger())
	model.View()
}
