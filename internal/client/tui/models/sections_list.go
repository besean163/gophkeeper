package models

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type SectionItem string

func (i SectionItem) FilterValue() string { return "" }

type SectionItemDelegate struct{}

func (d SectionItemDelegate) Height() int                             { return 1 }
func (d SectionItemDelegate) Spacing() int                            { return 0 }
func (d SectionItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d SectionItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SectionItem)
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

type SectionsModel struct {
	list list.Model
}

func NewSectionModel() *SectionsModel {
	items := []list.Item{
		SectionItem("Аккаунты"),
		SectionItem("Заметки"),
		SectionItem("Карты"),
	}

	l := list.New(items, SectionItemDelegate{}, 20, 10)
	l.Title = "Выберите секцию"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().MarginLeft(2)
	l.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	l.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	return &SectionsModel{
		list: l,
	}
}

func (m *SectionsModel) Init() tea.Cmd {
	return nil
}

func (m *SectionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *SectionsModel) View() string {
	return m.list.View()
}
