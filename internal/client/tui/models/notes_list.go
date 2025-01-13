package models

import (
	"fmt"
	"strings"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	keybinding "github.com/besean163/gophkeeper/internal/client/tui/models/key_binding"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NoteItem struct {
	item models.Note
}

func (i NoteItem) FilterValue() string { return i.item.Name }
func (i NoteItem) Title() string       { return i.item.Name }
func (i NoteItem) Description() string {
	desc := i.item.Content
	if len(i.item.Content) > 15 {
		desc = fmt.Sprintf("%s ...", string([]rune(i.item.Content)[0:10]))
	}
	return desc
}

// NoteListModel модель окна списка заметок
type NoteListModel struct {
	logger logger.Logger
	core   interfaces.Core
	list   list.Model
}

func NewNoteListModel(core interfaces.Core, logger logger.Logger) *NoteListModel {
	item := &NoteListModel{
		logger: logger,
		core:   core,
	}
	item.fiilList()
	return item
}

func (m *NoteListModel) Init() tea.Cmd {
	return nil
}

func (m *NoteListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Enter:
			if item, ok := m.list.SelectedItem().(NoteItem); ok {
				return m, func() tea.Msg { return messages.NoteEditMsg{Note: item.item} }
			}

			return m, func() tea.Msg { return messages.NoteEditMsg{} }
		case keybinding.Ctrlb:
			return m, func() tea.Msg { return messages.SectionBackMsg{} }
		case keybinding.Ctrla:
			return m, func() tea.Msg { return messages.NoteEditMsg{Note: models.Note{}} }
		case keybinding.Ctrld:
			if item, ok := m.list.SelectedItem().(NoteItem); ok {
				return m, func() tea.Msg { return messages.NoteDeleteMsg{Note: item.item} }
			}
		}
	case messages.NoteDeleteMsg:
		return m, m.delete(msg)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *NoteListModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Выбор заметки"))
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

func (m *NoteListModel) fiilList() {
	items, err := m.core.GetNotes()
	if err != nil {
		m.logger.Debug("account fill error", logger.NewField("error", err.Error()))
		items = make([]models.Note, 0)
	}
	listItems := make([]list.Item, 0)
	for _, item := range items {
		listItem := NoteItem{
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

func (m *NoteListModel) delete(msg messages.NoteDeleteMsg) tea.Cmd {
	err := m.core.DeleteNote(msg.Note)
	if err != nil {
		m.logger.Debug("note delete error", logger.NewField("error", err.Error()))
	}
	m.fiilList()
	return func() tea.Msg { return struct{}{} }
}
