package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextInputModel struct {
	textinput.Model
	style      lipgloss.Style
	focusStyle lipgloss.Style
}

func NewTextInputModel() TextInputModel {
	item := TextInputModel{
		Model: textinput.New(),
	}

	item.CharLimit = 32
	item.Cursor.SetMode(cursor.CursorBlink)

	return item
}

func (m *TextInputModel) WithFocusStyle(style lipgloss.Style) {
	m.focusStyle = style
}

func (m TextInputModel) Update(msg tea.Msg) (TextInputModel, tea.Cmd) {
	if m.Focused() {
		m.Cursor.Style = m.focusStyle
		m.TextStyle = m.focusStyle
	} else {
		m.Cursor.Style = m.style
		m.TextStyle = m.style
	}

	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	m.Prompt = fmt.Sprintf("%s: %s ", m.Prompt, m.getFocusPointer())
	result := m.Model.View()
	return result
}

func (m TextInputModel) getFocusPointer() string {
	if !m.Focused() {
		return " "
	}
	return m.focusStyle.Render(">")
}
