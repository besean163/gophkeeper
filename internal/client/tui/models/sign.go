package models

import (
	"fmt"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	signOptionSignIn = iota
	signOptionSignUp
)

var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	// progressEmpty = subtleStyle.Render(progressEmptyChar)
	dotStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(".")
	mainStyle = lipgloss.NewStyle().MarginLeft(2)
)

type SignModel struct {
	Choice int8
	Chosen bool
}

func NewSignModel() *SignModel {
	return &SignModel{
		Choice: signOptionSignIn,
		Chosen: false,
	}
}

func (m *SignModel) Init() tea.Cmd {
	return nil
}

func (m *SignModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("sign update")
	if !m.Chosen {
		return updateChoices(msg, m)
	}
	return m, nil
}

func (m *SignModel) View() string {
	logger.Get().Println("sign view")
	logger.Get().Println("view", m.Choice)

	c := m.Choice

	tpl := "Выберите способ входа:\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n",
		checkbox("Вход", c == signOptionSignIn),
		checkbox("Регистрация", c == signOptionSignUp),
	)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func updateChoices(msg tea.Msg, m *SignModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			m.Choice++
			if m.Choice > 1 {
				m.Choice = 1
			}
		case "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
		}
	}

	return m, nil
}
