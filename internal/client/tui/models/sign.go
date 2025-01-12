package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	keybinding "github.com/besean163/gophkeeper/internal/client/tui/models/key_binding"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	SignOptionLogin = iota
	SignOptionRegistration
)

// SignModel модель окна приветствия
type SignModel struct {
	logger         logger.Logger
	OptionSelected int
	OptionOrder    []int
	Options        []*components.OptionModel
}

func NewSignModel(logger logger.Logger) *SignModel {
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
		logger:      logger,
		OptionOrder: []int{SignOptionLogin, SignOptionRegistration},
		Options:     options,
	}
}

func (m *SignModel) Init() tea.Cmd {
	return nil
}

func (m *SignModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case messages.SignLoginMsg:
		m.logger.Debug("login")
		return m, nil
	case messages.SignRegistrationMsg:
		m.logger.Debug("registration")
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

	m.logger.Debug(b.String())
	return b.String()
}

func (m *SignModel) updateOptions(msg tea.Msg) tea.Cmd {
	option := m.Options[m.OptionSelected]
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Down:
			option.UnSelect()
			m.OptionSelected++
			if m.OptionSelected >= len(m.Options) {
				m.OptionSelected = len(m.Options) - 1
			}
			m.Options[m.OptionSelected].Select()

		case keybinding.Up:
			option.UnSelect()
			m.OptionSelected--
			if m.OptionSelected < 0 {
				m.OptionSelected = 0
			}
			m.Options[m.OptionSelected].Select()
		case keybinding.Enter:
			return m.Options[m.OptionSelected].Submit()
		}
	}
	return nil
}
