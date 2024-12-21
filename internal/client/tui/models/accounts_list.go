package models

import (
	"fmt"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
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

// type AccountItemDelegate struct{}

// func (d AccountItemDelegate) Height() int                             { return 1 }
// func (d AccountItemDelegate) Spacing() int                            { return 0 }
// func (d AccountItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
// func (d AccountItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
// 	i, ok := listItem.(AccountItem)
// 	if !ok {
// 		return
// 	}

// 	str := fmt.Sprintf("%d. %s", index+1, i)

// 	fn := lipgloss.NewStyle().PaddingLeft(4).Render
// 	if index == m.Index() {
// 		fn = func(s ...string) string {
// 			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + strings.Join(s, " "))
// 		}
// 	}

// 	fmt.Fprint(w, fn(str))
// }

type AccountListModel struct {
	list list.Model
}

func NewAccountListModel() *AccountListModel {
	item := &AccountListModel{}
	item.fiilList()
	return item
}

func (m *AccountListModel) Init() tea.Cmd {
	logger.Get().Println("init")

	return nil
}

func (m *AccountListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("update")

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			logger.Get().Println("choose msg")
			if item, ok := m.list.SelectedItem().(AccountItem); ok {
				logger.Get().Println("account name", item.account.Name)
			}

			// return m, func() tea.Msg { return messages.SectionBackMsg{} }
		case "ctrl+b":
			logger.Get().Println("back msg")
			return m, func() tea.Msg { return messages.SectionBackMsg{} }
		}
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
	return b.String()
}

func (m *AccountListModel) fiilList() {

	accounts := core.Instance.GetAccounts()
	items := make([]list.Item, 0)
	for _, account := range accounts {
		accountItem := AccountItem{
			account: account,
		}
		items = append(items, accountItem)
	}
	// logger.Get().Println(len(items))

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(styles.ColorAzure).
		Foreground(styles.ColorAzure).
		Padding(0, 0, 0, 1)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle

	l := list.New(items, delegate, 20, 15)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.FilterInput.Prompt = "Фильтр: "
	m.list = l
}
