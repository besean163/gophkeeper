package models

import tea "github.com/charmbracelet/bubbletea"

type NotesModel struct {
}

func (m *NotesModel) Init() tea.Cmd {
	return nil
}

func (m *NotesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *NotesModel) View() string {
	return ""
}
