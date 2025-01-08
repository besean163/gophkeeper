package models

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	tea "github.com/charmbracelet/bubbletea"
)

func TestSelectionListModel_init(t *testing.T) {
	model := models.NewSectionListModel(defaultlogger.NewDefaultLogger())
	model.Init()
}

func TestSelectionListModel_update(t *testing.T) {

	model := models.NewSectionListModel(defaultlogger.NewDefaultLogger())

	tests := []struct {
		name string
		msg  tea.Msg
	}{
		{
			name: "select accounts",
			msg:  messages.SelectAccountMsg{},
		},
		{
			name: "select notes",
			msg:  messages.SelectAccountMsg{},
		},
		{
			name: "select cards",
			msg:  messages.SelectAccountMsg{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model.Update(test.msg)
		})
	}
}

func TestSelectionListModel_view(t *testing.T) {
	model := models.NewSectionListModel(defaultlogger.NewDefaultLogger())
	model.View()
}
