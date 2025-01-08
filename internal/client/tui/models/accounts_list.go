package models

import (
	"fmt"
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

type AccountItem struct {
	account models.Account
}

func (i AccountItem) FilterValue() string { return i.account.Name }
func (i AccountItem) Title() string       { return i.account.Name }
func (i AccountItem) Description() string {
	return fmt.Sprintf("%s %s", i.account.Login, i.account.Password)
}

// AccountListModel модель окна списка аккаунтов
type AccountListModel struct {
	logger logger.Logger
	core   interfaces.Core
	list   list.Model
}

func NewAccountListModel(core interfaces.Core, logger logger.Logger) *AccountListModel {
	item := &AccountListModel{
		logger: logger,
		core:   core,
	}
	item.fiilList()
	return item
}

func (m *AccountListModel) Init() tea.Cmd {
	return nil
}

func (m *AccountListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.list.SelectedItem().(AccountItem); ok {
				return m, func() tea.Msg { return messages.AccountEditMsg{Account: item.account} }
			}

			return m, func() tea.Msg { return messages.AccountEditMsg{} }
		case "ctrl+b":
			return m, func() tea.Msg { return messages.SectionBackMsg{} }
		case "ctrl+a":
			return m, func() tea.Msg { return messages.AccountEditMsg{Account: models.Account{}} }
		case "ctrl+d":
			if item, ok := m.list.SelectedItem().(AccountItem); ok {
				return m, func() tea.Msg { return messages.AccountDeleteMsg{Account: item.account} }
			}
		}
	case messages.AccountDeleteMsg:
		return m, m.delete(msg)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *AccountListModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Выбор аккаунта"))
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

func (m *AccountListModel) fiilList() {

	accounts, err := m.core.GetAccounts()
	if err != nil {
		m.logger.Debug("account fill error", logger.NewField("error", err.Error()))
		accounts = make([]models.Account, 0)
	}
	items := make([]list.Item, 0)
	for _, account := range accounts {
		accountItem := AccountItem{
			account: account,
		}
		items = append(items, accountItem)
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(styles.ColorAzure).
		Foreground(styles.ColorAzure).
		Padding(0, 0, 0, 1)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle

	l := list.New(items, delegate, 20, 15)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.FilterInput.Prompt = "Фильтр: "
	m.list = l
}

func (m *AccountListModel) delete(msg messages.AccountDeleteMsg) tea.Cmd {
	err := m.core.DeleteAccount(msg.Account)
	if err != nil {
		m.logger.Debug("account delete error", logger.NewField("error", err.Error()))
	}
	m.fiilList()
	return func() tea.Msg { return struct{}{} }
}
