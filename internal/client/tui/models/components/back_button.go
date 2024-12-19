package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BackButtonModel struct {
	name     string
	focus    bool
	pressKey string
}

type BackButtonPressMsg struct{}

func NewBackButtonModel(name string) *BackButtonModel {
	return &BackButtonModel{
		name: name,
	}
}

func (m *BackButtonModel) WithKey(key string) *BackButtonModel {
	m.pressKey = key
	return m
}

func (m *BackButtonModel) Init() tea.Cmd {
	return nil
}

func (m *BackButtonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.pressKey == "" {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == m.pressKey {
			return m, func() tea.Msg { return BackButtonPressMsg{} }
		}
	}
	return m, nil
}

func (m *BackButtonModel) View() string {
	button := fmt.Sprintf("[ %s ]", m.name)
	if m.focus {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(button)
	}
	return button
}

func (m *BackButtonModel) Focus() {
	m.focus = true
}

func (m *BackButtonModel) Blur() {
	m.focus = false
}
