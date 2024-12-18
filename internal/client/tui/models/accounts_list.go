package models

import (
	"fmt"
	"io"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AccountItem string

func (i AccountItem) FilterValue() string { return "" }

type AccountItemDelegate struct{}

func (d AccountItemDelegate) Height() int                             { return 1 }
func (d AccountItemDelegate) Spacing() int                            { return 0 }
func (d AccountItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d AccountItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(AccountItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := lipgloss.NewStyle().PaddingLeft(4).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

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
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *AccountListModel) View() string {
	return m.list.View()
}

func (m *AccountListModel) fiilList() {

	accounts := core.Instance.GetAccounts()
	items := make([]list.Item, 0)
	for _, account := range accounts {
		items = append(items, AccountItem(account.Name))
	}
	// logger.Get().Println(len(items))

	l := list.New(items, AccountItemDelegate{}, 20, 10)
	m.list = l
}
