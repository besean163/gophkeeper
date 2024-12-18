package models

import tea "github.com/charmbracelet/bubbletea"

type CartListModel struct {
}

func (m *CartListModel) Init() tea.Cmd {
	return nil
}

func (m *CartListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *CartListModel) View() string {
	return ""
}
