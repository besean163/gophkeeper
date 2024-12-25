package app

import (
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	tea "github.com/charmbracelet/bubbletea"
)

func NewProgram() *tea.Program {
	return tea.NewProgram(models.NewRootModel())
}
