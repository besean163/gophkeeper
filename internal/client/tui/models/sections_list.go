package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	SectionOptionAccount = iota
	SectionOptionNote
	SectionOptionCard
)

type SectionListModel struct {
	OptionSelected int
	OptionOrder    []int
	Options        []*components.OptionModel
}

func NewSectionListModel() *SectionListModel {
	options := []*components.OptionModel{
		components.NewOption("1. Аккаунты").
			WithSelectedName("1. Аккаунты \u279C").
			WithSubmitMsg(messages.SelectAccountMsg{}).
			WithStyle(lipgloss.NewStyle().PaddingLeft(4)).
			WithSelectStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure)).
			Select(),
		components.NewOption("2. Заметки").
			WithSelectedName("2. Заметки \u279C").
			WithSubmitMsg(messages.SelectNoteMsg{}).
			WithStyle(lipgloss.NewStyle().PaddingLeft(4)).
			WithSelectStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure)),
		components.NewOption("3. Карты").
			WithSelectedName("3. Карты \u279C").
			WithSubmitMsg(messages.SelectCardMsg{}).
			WithStyle(lipgloss.NewStyle().PaddingLeft(4)).
			WithSelectStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure)),
	}
	return &SectionListModel{
		OptionOrder: []int{SectionOptionAccount, SectionOptionNote, SectionOptionCard},
		Options:     options,
	}
}

func (m *SectionListModel) Init() tea.Cmd {
	return nil
}

func (m *SectionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("update")
	switch msg.(type) {
	case messages.SelectAccountMsg:
		logger.Get().Println("choose accounts")
		return m, nil
	case messages.SelectNoteMsg:
		logger.Get().Println("choose notes")
		return m, nil
	case messages.SelectCardMsg:
		logger.Get().Println("choose carts")
		return m, nil
	}

	return m, m.updateOptions(msg)
}

func (m *SectionListModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Выбор раздела"))
	b.WriteByte('\n')
	b.WriteByte('\n')

	for _, option := range m.Options {
		b.WriteString(option.View())
		b.WriteByte('\n')
	}

	return b.String()
}

func (m *SectionListModel) updateOptions(msg tea.Msg) tea.Cmd {
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
