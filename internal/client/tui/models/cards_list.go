package models

import (
	"strconv"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CardItem struct {
	item models.Card
}

func (i CardItem) FilterValue() string { return i.item.Name }
func (i CardItem) Title() string       { return i.item.Name }
func (i CardItem) Description() string {
	return strconv.Itoa(i.item.Number)
}

// CardListModel модель окна списка карт
type CardListModel struct {
	logger logger.Logger
	core   interfaces.Core
	list   list.Model
}

func NewCardListModel(core interfaces.Core, logger logger.Logger) *CardListModel {
	item := &CardListModel{
		logger: logger,
		core:   core,
	}
	item.fiilList()
	return item
}

func (m *CardListModel) Init() tea.Cmd { return nil }

func (m *CardListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.list.SelectedItem().(CardItem); ok {
				return m, func() tea.Msg { return messages.CardEditMsg{Card: item.item} }
			}

			return m, func() tea.Msg { return messages.CardEditMsg{} }
		case "ctrl+b":
			return m, func() tea.Msg { return messages.SectionBackMsg{} }
		case "ctrl+a":
			return m, func() tea.Msg { return messages.CardEditMsg{Card: models.Card{}} }
		case "ctrl+d":
			if item, ok := m.list.SelectedItem().(CardItem); ok {
				return m, func() tea.Msg { return messages.CardDeleteMsg{Card: item.item} }
			}
		}
	case messages.CardDeleteMsg:
		return m, m.delete(msg)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *CardListModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Выбор карты"))
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render(m.list.View()))
	b.WriteRune('\n')
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+a: add"))
	b.WriteRune('\n')
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+d: delete"))
	b.WriteRune('\n')
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(" enter: edit"))
	return b.String()
}

func (m *CardListModel) fiilList() {
	items, err := m.core.GetCards()
	if err != nil {
		m.logger.Debug("card fill error", logger.NewField("error", err.Error()))
		items = make([]models.Card, 0)
	}
	listItems := make([]list.Item, 0)
	for _, item := range items {
		listItem := CardItem{
			item: item,
		}
		listItems = append(listItems, listItem)
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(styles.ColorAzure).
		Foreground(styles.ColorAzure).
		Padding(0, 0, 0, 1)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle

	l := list.New(listItems, delegate, 20, 15)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.FilterInput.Prompt = "Фильтр: "
	m.list = l
}

func (m *CardListModel) delete(msg messages.CardDeleteMsg) tea.Cmd {
	err := m.core.DeleteCard(msg.Card)
	if err != nil {
		m.logger.Debug("card delete error", logger.NewField("error", err.Error()))
	}
	m.fiilList()
	return func() tea.Msg { return struct{}{} }
}
