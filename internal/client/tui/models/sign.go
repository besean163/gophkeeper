package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	SignOptionLogin = iota
	SignOptionRegistration
)

type SignModel struct {
	OptionSelected int
	OptionOrder    []int
	Options        []*components.OptionModel
}

func NewSignModel() *SignModel {
	options := []*components.OptionModel{
		components.NewOption("Вход").
			WithSelectedName("Вход \u279C").
			WithSubmitMsg(messages.SignLoginMsg{}).
			WithStyle(lipgloss.NewStyle().PaddingLeft(4)).
			WithSelectStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure)).
			Select(),
		components.NewOption("Регистрация").
			WithSelectedName("Регистрация \u279C").
			WithSubmitMsg(messages.SignRegistrationMsg{}).
			WithStyle(lipgloss.NewStyle().PaddingLeft(4)).
			WithSelectStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGreen)),
	}
	return &SignModel{
		OptionOrder: []int{SignOptionLogin, SignOptionRegistration},
		Options:     options,
	}
}

func (m *SignModel) Init() tea.Cmd {
	return nil
}

func (m *SignModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("update")
	switch msg.(type) {
	case messages.SignLoginMsg:
		logger.Get().Println("login")
		return m, nil
	case messages.SignRegistrationMsg:
		logger.Get().Println("registration")
		return m, nil
	}

	return m, m.updateOptions(msg)
}

func (m *SignModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Gophkeeper v1.0"))
	b.WriteByte('\n')
	b.WriteByte('\n')

	for _, option := range m.Options {
		b.WriteString(option.View())
		b.WriteByte('\n')
	}

	return b.String()
}

func (m *SignModel) updateOptions(msg tea.Msg) tea.Cmd {
	option := m.Options[m.OptionSelected]
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			option.UnSelect()
			m.OptionSelected++
			if m.OptionSelected >= len(m.Options) {
				m.OptionSelected = len(m.Options) - 1
			}
			m.Options[m.OptionSelected].Select()

		case "up":
			option.UnSelect()
			m.OptionSelected--
			if m.OptionSelected < 0 {
				m.OptionSelected = 0
			}
			m.Options[m.OptionSelected].Select()
		case "enter":
			return m.Options[m.OptionSelected].Submit()
		}
	}
	return nil
}
