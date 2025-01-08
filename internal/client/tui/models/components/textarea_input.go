package components

import (
	"fmt"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextAreaInputModel struct {
	Name string
	textarea.Model
	style      lipgloss.Style
	focusStyle lipgloss.Style
}

func NewTextAreaInputModel() TextAreaInputModel {
	item := TextAreaInputModel{
		Model: textarea.New(),
	}

	item.ShowLineNumbers = false
	item.Placeholder = "введите текст ..."

	return item
}

func (m *TextAreaInputModel) WithFocusStyle(style lipgloss.Style) {
	m.Model.FocusedStyle.CursorLine = lipgloss.NewStyle().Foreground(styles.ColorAzure)
	m.focusStyle = style
}

func (m TextAreaInputModel) Update(msg tea.Msg) (TextAreaInputModel, tea.Cmd) {
	if m.Focused() {
		m.Cursor.Style = m.focusStyle
	} else {
		m.Cursor.Style = m.style
	}

	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m TextAreaInputModel) View() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("%s:", m.Name))
	b.WriteRune('\n')
	b.WriteString(m.Model.View())
	return b.String()
}
