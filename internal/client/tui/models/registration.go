package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	RegistrationGroupInput = iota
	RegistrationGroupButtons
)

const (
	RegistrationInputLogin = iota
	RegistrationInputPassword
	RegistrationInputPasswordRepeat
)

const (
	RegistrationButtonEnter = iota
	RegistrationButtonBack
)

// RegistrationModel модель окна регистрации
type RegistrationModel struct {
	logger  logger.Logger
	core    interfaces.Core
	fc      *components.GroupFocusCursor
	inputs  []components.TextInputModel
	buttons []components.ButtonModel
}

func NewRegistrationModel(core interfaces.Core, logger logger.Logger) *RegistrationModel {
	item := &RegistrationModel{
		logger: logger,
		core:   core,
		fc:     components.NewGroupFocusCursor(RegistrationGroupInput, 0),
	}
	item.setInputs()
	item.setButtons()
	item.activateButtons()
	return item
}

func (m *RegistrationModel) Init() tea.Cmd {
	return nil
}

func (m *RegistrationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.activateButtons()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			var cmd tea.Cmd
			cmd = m.pressButtons(msg)
			if cmd != nil {
				return m, cmd
			}
			m.moveFocus(msg)
		}
	case messages.ButtonSubmitMsg:
		return m, m.registration()
	case messages.ButtonBackMsg:
		return m, func() tea.Msg { return messages.SignBackMsg{} }
	}
	m.updateInputs(msg)

	return m, nil
}

func (m *RegistrationModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Вход"))
	b.WriteRune('\n')
	b.WriteRune('\n')
	for _, input := range m.inputs {
		b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(input.View()))
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for _, button := range m.buttons {
		b.WriteString(button.View())
		b.WriteRune('\n')
	}
	return b.String()
}

func (m *RegistrationModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

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
		if key == "up" || key == "shift+tab" {
			m.fc.Index--
		} else {
			m.fc.Index++
		}

		// переключить на другой список если требуется и установить верный индекс
		if m.fc.Group == RegistrationGroupInput && m.fc.Index >= len(m.inputs) {
			m.fc.Move(RegistrationGroupButtons, 0)
		}

		if m.fc.Group == RegistrationGroupButtons && (m.fc.Index >= len(m.buttons) || m.fc.Index < 0) {
			m.fc.Move(RegistrationGroupInput, 0)
		}

		if m.fc.Group == RegistrationGroupInput {
			if m.fc.Index < 0 {
				m.fc.Index = 0
			}
			return m.inputs[m.fc.Index].Focus()
		}

		if m.fc.Group == RegistrationGroupButtons {
			// если кнопка активна оставляем фокус
			if m.buttons[m.fc.Index].IsActive() {
				m.buttons[m.fc.Index].Focus()
				break
			}

		}
	}

	return nil

}

func (m *RegistrationModel) updateInputs(msg tea.Msg) tea.Cmd {
	// logger.Get().Println("updates")
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *RegistrationModel) activateButtons() {
	m.buttons[RegistrationButtonEnter].Activate(func() bool { return m.isValid() })
	m.buttons[RegistrationButtonBack].Activate(func() bool { return true })
}

func (m *RegistrationModel) pressButtons(msg tea.Msg) tea.Cmd {
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

func (m *RegistrationModel) isValid() bool {

	login := m.inputs[RegistrationInputLogin].Value()
	password := m.inputs[RegistrationInputPassword].Value()

	return login != "" && password != ""
}

func (m *RegistrationModel) setInputs() {

	var input components.TextInputModel
	items := make([]components.TextInputModel, 3)
	input = components.NewTextInputModel()
	input.Focus()
	input.Prompt = "        Логин"
	input.Placeholder = "введите логин ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	items[RegistrationInputLogin] = input

	input = components.NewTextInputModel()
	input.Prompt = "       Пароль"
	input.Placeholder = "введите пароль ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	items[RegistrationInputPassword] = input

	input = components.NewTextInputModel()
	input.Prompt = "Повтор пароля"
	input.Placeholder = "повторите пароль ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	items[RegistrationInputPasswordRepeat] = input

	m.inputs = items
}

func (m *RegistrationModel) setButtons() {
	var button components.ButtonModel
	items := make([]components.ButtonModel, 2)

	// добавляем кнопку вход
	button = components.NewButtonModel("Регистрация")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGreen))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonSubmitMsg{})
	items[RegistrationButtonEnter] = button

	// добавляем кнопку назад
	button = components.NewButtonModel("Назад")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorOrange))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonBackMsg{})
	items[RegistrationButtonBack] = button

	m.buttons = items
}

func (m *RegistrationModel) registration() tea.Cmd {
	err := m.core.Register(m.inputs[LoginInputLogin].Value(), m.inputs[LoginInputPassword].Value())

	if err != nil {
		m.logger.Debug("registration error", logger.NewField("error", err.Error()))
		return nil
	} else {
		m.logger.Debug("registration success continue")
	}

	return func() tea.Msg { return messages.RegistrationSuccessMsg{} }
}
