package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ErrorMessageModel struct {
	Show    bool
	Message string
}

func NewErrorMessageModel() *ErrorMessageModel {
	return &ErrorMessageModel{}
}

func (m *ErrorMessageModel) Init() tea.Cmd {
	return nil
}

func (m *ErrorMessageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *ErrorMessageModel) View() string {
	if !m.Show {
		return ""
	}

	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).BorderForeground(lipgloss.Color("#FF5A40")).Render(m.Message)
}
