package models

import tea "github.com/charmbracelet/bubbletea"

type SignModel struct {
}

func (m SignModel) Init() tea.Cmd {
	return nil
}

func (m SignModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SignModel) View() string {
	return ""
}
