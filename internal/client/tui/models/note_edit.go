package models

import (
	"strings"

	coremodels "github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	NoteEditGroupInput = iota
	NoteEditGroupTextArea
	NoteEditGroupButtons
)

const (
	NoteEditButtonSave = iota
	NoteEditButtonBack
)

// NoteEditModel модель окна редактирования заметки
type NoteEditModel struct {
	logger   logger.Logger
	core     interfaces.Core
	item     coremodels.Note
	fc       *components.GroupFocusCursor
	input    components.TextInputModel
	textarea components.TextAreaInputModel
	buttons  []components.ButtonModel
}

func NewNoteEditModel(core interfaces.Core, item coremodels.Note, logger logger.Logger) *NoteEditModel {
	model := &NoteEditModel{
		logger: logger,
		core:   core,
		fc:     components.NewGroupFocusCursor(NoteEditGroupInput, 0),
		item:   item,
	}

	model.setInputs()
	model.setTextArea()
	model.setButtons()
	return model
}

func (m *NoteEditModel) Init() tea.Cmd {
	return nil
}

func (m *NoteEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	m.activateButtons()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab":
			m.moveFocus(msg)
		case "enter":
			var cmd tea.Cmd
			cmd = m.pressButtons(msg)
			if cmd != nil {
				return m, cmd
			}
		}
	case messages.ButtonSubmitMsg:
		m.logger.Info("save msg")
		return m, m.save()
	case messages.ButtonBackMsg:
		m.logger.Info("back msg")
		return m, func() tea.Msg { return messages.NoteListBackMsg{} }
	}
	m.updateInputs(msg)
	m.activateButtons()

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *NoteEditModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Редактирование заметки"))
	b.WriteRune('\n')
	b.WriteRune('\n')

	b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(m.input.View()))
	b.WriteRune('\n')
	b.WriteRune('\n')

	b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(m.textarea.View()))
	b.WriteRune('\n')
	b.WriteRune('\n')

	b.WriteRune('\n')
	for _, button := range m.buttons {
		b.WriteString(button.View())
		b.WriteByte('\n')
	}
	return b.String()
}

func (m *NoteEditModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

	// зачищаем фокус
	m.input.Blur()
	m.textarea.Blur()

	// проставляем активность на кнопках, чтобы не активные не выбирать
	m.activateButtons()
	for i := range m.buttons {
		m.buttons[i].Blur()
	}

	for {
		key := msg.String()
		if key == "shift+tab" {
			m.fc.Index--
		} else {
			m.fc.Index++
		}

		if m.fc.Group == NoteEditGroupInput {
			if m.fc.Index >= 1 {
				m.fc.Move(NoteEditGroupTextArea, 0)
			} else if m.fc.Index < 0 {
				m.fc.Move(NoteEditGroupButtons, len(m.buttons)-1)
			}
		}

		if m.fc.Group == NoteEditGroupTextArea {
			if m.fc.Index >= 1 {
				m.fc.Move(NoteEditGroupButtons, 0)
			} else if m.fc.Index < 0 {
				m.fc.Move(NoteEditGroupInput, 0)
			}
		}

		if m.fc.Group == NoteEditGroupButtons {
			if m.fc.Index >= len(m.buttons) {
				m.fc.Move(NoteEditGroupInput, 0)
			} else if m.fc.Index < 0 {
				m.fc.Move(NoteEditGroupTextArea, 0)
			}
		}

		if m.fc.Group == NoteEditGroupInput {
			return m.input.Focus()
		}

		if m.fc.Group == NoteEditGroupTextArea {
			return m.textarea.Focus()
		}

		if m.fc.Group == NoteEditGroupButtons {
			// если кнопка активна оставляем фокус
			if m.buttons[m.fc.Index].IsActive() {
				m.buttons[m.fc.Index].Focus()
				break
			}

		}
	}

	return nil
}

func (m *NoteEditModel) activateButtons() {
	m.buttons[AccountEditButtonSave].Activate(func() bool { return m.isValid() })
	m.buttons[AccountEditButtonBack].Activate(func() bool { return true })
}

func (m *NoteEditModel) pressButtons(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			var cmd tea.Cmd
			for _, button := range m.buttons {
				cmd = button.Press()
				if cmd != nil {
					return cmd
				}
			}
		}
	}
	return nil
}

func (m *NoteEditModel) isValid() bool {
	name := m.input.Value()
	return name != ""
}

func (m *NoteEditModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 2)

	m.input, cmds[0] = m.input.Update(msg)
	m.textarea, cmds[1] = m.textarea.Update(msg)

	return tea.Batch(cmds...)
}

func (m *NoteEditModel) setInputs() {

	input := components.NewTextInputModel()
	input.Focus()
	input.Prompt = "Название"
	input.Placeholder = "введите название ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.item.Name)

	m.input = input
}

func (m *NoteEditModel) setTextArea() {

	input := components.NewTextAreaInputModel()
	input.Name = "Текст"
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.item.Content)

	m.textarea = input
}

func (m *NoteEditModel) setButtons() {
	var button components.ButtonModel
	items := make([]components.ButtonModel, 2)

	// добавляем кнопку вход
	button = components.NewButtonModel("Сохранить")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGreen))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonSubmitMsg{})
	items[AccountEditButtonSave] = button

	// добавляем кнопку назад
	button = components.NewButtonModel("Отмена")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorRed))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonBackMsg{})
	items[AccountEditButtonBack] = button

	m.buttons = items
}

func (m *NoteEditModel) save() tea.Cmd {
	item := m.item
	item.ID = m.item.ID
	item.Name = m.input.Value()
	item.Content = m.textarea.Value()

	err := m.core.SaveNote(item)

	if err != nil {
		m.logger.Info("error")
	} else {
		m.logger.Info("save success continue")
	}

	return func() tea.Msg { return messages.NoteListBackMsg{} }
}
