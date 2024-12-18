package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EnterButtonModel struct {
	name     string
	focus    bool
	pressKey string
}

type EnterButtonPressMsg struct{}

func NewEnterButtonModel(name string) *EnterButtonModel {
	return &EnterButtonModel{
		name: name,
	}
}

func (m *EnterButtonModel) WithKey(key string) *EnterButtonModel {
	m.pressKey = key
	return m
}

func (m *EnterButtonModel) Init() tea.Cmd {
	return nil
}

func (m *EnterButtonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.pressKey == "" {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == m.pressKey {
			return m, func() tea.Msg { return EnterButtonPressMsg{} }
		}
	}
	return m, nil
}

func (m *EnterButtonModel) View() string {
	button := fmt.Sprintf("[ %s ]", m.name)
	if m.focus {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(button)
	}
	return button
}

func (m *EnterButtonModel) Focus() {
	m.focus = true
}

func (m *EnterButtonModel) Blur() {
	m.focus = false
}
