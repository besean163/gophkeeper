package models

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	tea "github.com/charmbracelet/bubbletea"
)

func TestSign_init(t *testing.T) {
	model := models.NewSignModel(defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestSign_update(t *testing.T) {
	model := models.NewSignModel(defaultlogger.NewDefaultLogger())

	tests := []struct {
		name string
		msg  tea.Msg
	}{
		{
			name: "simple",
			msg:  struct{}{},
		},
		{
			name: "choose login",
			msg:  messages.SignLoginMsg{},
		},
		{
			name: "choose registration",
			msg:  messages.SignRegistrationMsg{},
		},
		{
			name: "down",
			msg:  tea.KeyMsg{Type: tea.KeyDown},
		},
		{
			name: "up",
			msg:  tea.KeyMsg{Type: tea.KeyUp},
		},
		{
			name: "enter",
			msg:  tea.KeyMsg{Type: tea.KeyEnter},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model.Update(test.msg)
		})
	}
}

func TestSign_view(t *testing.T) {
	model := models.NewSignModel(defaultlogger.NewDefaultLogger())
	model.View()
}
