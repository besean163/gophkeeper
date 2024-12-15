package models

import tea "github.com/charmbracelet/bubbletea"

type SectionsModel struct {
}

func (m SectionsModel) Init() tea.Cmd {
	return nil
}

func (m SectionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SectionsModel) View() string {
	return ""
}
