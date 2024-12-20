package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ButtonModel struct {
	name           string
	focus          bool
	active         bool
	style          lipgloss.Style
	focusStyle     lipgloss.Style
	notActiveStyle lipgloss.Style
	pressMsg       tea.Msg
}

func NewButtonModel(name string) ButtonModel {
	return ButtonModel{
		name:   name,
		active: true,
	}
}

func (b *ButtonModel) WithStyle(style lipgloss.Style) *ButtonModel {
	b.style = style
	return b
}

func (b *ButtonModel) WithSelectedStyle(style lipgloss.Style) *ButtonModel {
	b.focusStyle = style
	return b
}

func (b *ButtonModel) WithNotActiveStyle(style lipgloss.Style) *ButtonModel {
	b.notActiveStyle = style
	return b
}

func (b *ButtonModel) IsActive() bool {
	return b.active
}

func (b *ButtonModel) Activate(fn func() bool) *ButtonModel {
	b.active = fn()
	return b
}

func (b *ButtonModel) Focus() {
	b.focus = true
}

func (b *ButtonModel) Blur() {
	b.focus = false
}

func (b *ButtonModel) Press() tea.Cmd {
	if b.pressMsg != nil {
		return func() tea.Msg { return b.pressMsg }
	}
	return nil
}

func (b *ButtonModel) View() string {
	result := fmt.Sprintf("[ %s ]", b.name)

	if !b.active {
		return b.notActiveStyle.Render(result)
	}

	if b.focus {
		return b.focusStyle.Render(result)
	}

	return b.style.Render(result)
}
