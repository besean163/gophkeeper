package models

import tea "github.com/charmbracelet/bubbletea"

type AccountsModel struct {
}

func (m AccountsModel) Init() tea.Cmd {
	return nil
}

func (m AccountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m AccountsModel) View() string {
	return ""
}
