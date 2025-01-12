package models

import (
	"strconv"
	"strings"

	coremodels "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	keybinding "github.com/besean163/gophkeeper/internal/client/tui/models/key_binding"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	CardEditGroupInput = iota
	CardEditGroupButtons
)

const (
	CardEditInputName = iota
	CardEditInputNumber
	CardEditInputExp
	CardEditInputCVV
)

const (
	CardEditButtonSave = iota
	CardEditButtonBack
)

// CardEditModel модель окна редактирования карты
type CardEditModel struct {
	logger  logger.Logger
	core    interfaces.Core
	item    coremodels.Card
	fc      *components.GroupFocusCursor
	inputs  []components.TextInputModel
	buttons []components.ButtonModel
}

func NewCardEditModel(core interfaces.Core, item coremodels.Card, logger logger.Logger) *CardEditModel {
	model := &CardEditModel{
		logger: logger,
		core:   core,
		fc:     components.NewGroupFocusCursor(CardEditGroupInput, 0),
		item:   item,
	}

	model.setInputs()
	model.setButtons()
	return model
}

func (m *CardEditModel) Init() tea.Cmd {
	return nil
}

func (m *CardEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == keybinding.Ctrlc {
			return m, tea.Quit
		}
	}

	m.activateButtons()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Tab, keybinding.ShiftTab, keybinding.Enter, keybinding.Up, keybinding.Down:
			var cmd tea.Cmd
			cmd = m.pressButtons(msg)
			if cmd != nil {
				return m, cmd
			}
			m.logger.Info("update")
			m.moveFocus(msg)
		}
	case messages.ButtonSubmitMsg:
		m.logger.Info("save msg")
		return m, m.save()
	case messages.ButtonBackMsg:
		m.logger.Info("back msg")
		return m, func() tea.Msg { return messages.CardListBackMsg{} }
	}
	m.updateInputs(msg)
	m.activateButtons()
	return m, nil
}

func (m *CardEditModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Редактирование карты"))
	b.WriteRune('\n')
	b.WriteRune('\n')
	for _, input := range m.inputs {
		b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(input.View()))
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for _, button := range m.buttons {
		b.WriteString(button.View())
		b.WriteByte('\n')
	}
	return b.String()
}

func (m *CardEditModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

	for i := range m.inputs {
		m.inputs[i].Blur()
	}

	// проставляем активность на кнопках, чтобы не активные не выбирать
	m.activateButtons()
	for i := range m.buttons {
		m.buttons[i].Blur()
	}

	for {
		key := msg.String()
		if key == keybinding.Up || key == keybinding.ShiftTab {
			m.fc.Index--
		} else {
			m.fc.Index++
		}

		// переключить на другой список если требуется и установить верный индекс
		if m.fc.Group == CardEditGroupInput && m.fc.Index >= len(m.inputs) {
			m.fc.Move(CardEditGroupButtons, 0)
		}

		if m.fc.Group == CardEditGroupButtons {
			if m.fc.Index >= len(m.buttons) {
				m.fc.Move(CardEditGroupInput, 0)
			} else if m.fc.Index < 0 {
				m.fc.Move(CardEditGroupInput, len(m.inputs)-1)
			}
		}

		if m.fc.Group == CardEditGroupInput {
			if m.fc.Index < 0 {
				m.fc.Index = 0
			}
			return m.inputs[m.fc.Index].Focus()
		}

		if m.fc.Group == CardEditGroupButtons {
			// если кнопка активна оставляем фокус
			if m.buttons[m.fc.Index].IsActive() {
				m.buttons[m.fc.Index].Focus()
				break
			}

		}
	}

	return nil
}

func (m *CardEditModel) activateButtons() {
	m.buttons[AccountEditButtonSave].Activate(func() bool { return m.isValid() })
	m.buttons[AccountEditButtonBack].Activate(func() bool { return true })
}

func (m *CardEditModel) pressButtons(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Enter:
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

func (m *CardEditModel) isValid() bool {
	name := m.inputs[CardEditInputName].Value()
	number := m.inputs[CardEditInputNumber].Value()
	exp := m.inputs[CardEditInputExp].Value()
	cvv := m.inputs[CardEditInputCVV].Value()

	return name != "" && number != "" && exp != "" && cvv != ""
}

func (m *CardEditModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *CardEditModel) setInputs() {

	var input components.TextInputModel
	items := make([]components.TextInputModel, 4)

	input = components.NewTextInputModel()
	input.Focus()
	input.Prompt = "Название"
	input.Placeholder = "введите название ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.item.Name)
	items[CardEditInputName] = input

	input = components.NewTextInputModel()
	input.Prompt = "   Номер"
	input.Placeholder = "введите номер ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(strconv.Itoa(m.item.Number))
	items[CardEditInputNumber] = input

	input = components.NewTextInputModel()
	input.Prompt = "Истекает"
	input.Placeholder = "введите MM/YY ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.item.Exp)
	items[CardEditInputExp] = input

	input = components.NewTextInputModel()
	input.Prompt = "     CVV"
	input.Placeholder = "введите CVV ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(strconv.Itoa(m.item.CVV))
	items[CardEditInputCVV] = input

	m.inputs = items
}

func (m *CardEditModel) setButtons() {
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

func (m *CardEditModel) save() tea.Cmd {
	var err error
	number, err := strconv.Atoi(m.inputs[CardEditInputNumber].Value())
	if err != nil {
		m.logger.Info(err.Error())
	}

	cvv, err := strconv.Atoi(m.inputs[CardEditInputCVV].Value())
	if err != nil {
		m.logger.Info(err.Error())
	}

	item := m.item
	item.Name = m.inputs[CardEditInputName].Value()
	item.Number = number
	item.Exp = m.inputs[CardEditInputExp].Value()
	item.CVV = cvv

	err = m.core.SaveCard(item)

	if err != nil {
		m.logger.Info("error")
	} else {
		m.logger.Info("save success continue")
	}

	return func() tea.Msg { return messages.CardListBackMsg{} }
}
