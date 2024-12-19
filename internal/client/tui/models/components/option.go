package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OptionModel struct {
	name         string
	selectedName string
	selected     bool
	submitted    bool
	style        *lipgloss.Style
	selectStyle  *lipgloss.Style
	submitMsg    tea.Msg
}

func NewOption(name string) *OptionModel {
	return &OptionModel{
		name: name,
	}
}

func (o *OptionModel) Select() *OptionModel {
	o.selected = true
	return o
}

func (o *OptionModel) UnSelect() *OptionModel {
	o.selected = false
	return o
}

func (o *OptionModel) WithSelectedName(name string) *OptionModel {
	o.selectedName = name
	return o
}

func (o *OptionModel) WithStyle(style lipgloss.Style) *OptionModel {
	o.style = &style
	return o
}

func (o *OptionModel) WithSelectStyle(style lipgloss.Style) *OptionModel {
	o.selectStyle = &style
	return o
}

func (o *OptionModel) WithSubmitMsg(msg tea.Msg) *OptionModel {
	o.submitMsg = msg
	return o
}

func (o *OptionModel) Submit() tea.Cmd {
	if o.submitMsg != nil {
		return func() tea.Msg { return o.submitMsg }
	}
	return nil
}

func (o *OptionModel) Submitted() bool {
	return o.submitted
}

func (m *OptionModel) View() string {
	result := m.name

	if m.selected && m.selectedName != "" {
		result = m.selectedName
	}

	if m.selected && m.selectStyle != nil {
		result = m.selectStyle.Render(result)
	}

	if !m.selected && m.style != nil {
		result = m.style.Render(result)
	}

	return result
}
