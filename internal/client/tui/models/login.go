package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/interfaces"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	inputLogin = iota
	inputPassword
	inputRepeatedPassword
)

const (
	focusGroupInput = iota
	focusGroupControl
)

type FocusCursor struct {
	Group int
	Index int
}

func (fc *FocusCursor) move(group, index int) {
	fc.Group = group
	fc.Index = index
}

type TypedTextInput struct {
	Type int
	textinput.Model
}

func (m TypedTextInput) Update(msg tea.Msg) (TypedTextInput, tea.Cmd) {
	i, cmd := m.Model.Update(msg)
	m.Model = i
	return m, cmd
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	// blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	noStyle = lipgloss.NewStyle()
	// focusedButton = focusedStyle.Render("[ Вход ]")
	// blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Вход"))

	// errorStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).BorderForeground(lipgloss.Color("#FF5A40"))
)

type LoginModel struct {
	fc            *FocusCursor
	inputs        []TypedTextInput
	controlInputs []interfaces.ControlButton
	registration  bool
	errorMessage  *components.ErrorMessageModel
}

func NewLoginModel(registration bool) *LoginModel {

	inputs := make([]TypedTextInput, 0)
	inputs = append(inputs, NewLoginTextInput())
	inputs = append(inputs, NewPasswordTextInput())
	if registration {
		inputs = append(inputs, NewRepeatPasswordTextInput())
	}

	controlInputs := make([]interfaces.ControlButton, 0)
	controlInputs = append(controlInputs, components.NewEnterButtonModel("Вход").WithKey("enter"))

	return &LoginModel{
		fc:            &FocusCursor{Group: focusGroupInput, Index: 0},
		registration:  registration,
		inputs:        inputs,
		controlInputs: controlInputs,
		errorMessage:  &components.ErrorMessageModel{},
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	cmds := make([]tea.Cmd, len(m.inputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			m.moveFocus(msg)
			cmds = append(cmds, m.updateFocusInputs())
		}
	case components.EnterButtonPressMsg:
		logger.Get().Println("press")
		if m.isValid() {
			return m.login()
		}
	}

	cmds = append(cmds, m.updateInputs(msg))

	return m, tea.Batch(cmds...)
}

func (m *LoginModel) View() string {
	var b strings.Builder
	b.WriteRune('\n')
	b.WriteString("Введите ваши данные:\n")

	// отображаем текстовые инпуты
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		b.WriteRune('\n')
	}

	b.WriteString(m.errorMessage.View())

	b.WriteRune('\n')

	// отображаем кнопки
	for i := range m.controlInputs {
		b.WriteString(m.controlInputs[i].View())
		b.WriteRune('\n')
	}

	return b.String()
}

func (m *LoginModel) moveFocus(msg tea.KeyMsg) {

	s := msg.String()
	if s == "up" || s == "shift+tab" {
		m.fc.Index--
	} else {
		m.fc.Index++
	}

	if m.fc.Group == focusGroupInput {
		if m.fc.Index < 0 {
			m.fc.move(focusGroupInput, 0)
		} else if m.fc.Index >= len(m.inputs) {
			m.fc.move(focusGroupControl, 0)
		}
	} else if m.fc.Group == focusGroupControl {
		if m.fc.Index < 0 {
			m.fc.move(focusGroupInput, len(m.inputs)-1)
		} else if m.fc.Index >= len(m.controlInputs) {
			m.fc.move(focusGroupControl, len(m.controlInputs)-1)
		}
	}
}

func (m *LoginModel) updateFocusInputs() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if m.fc.Group == focusGroupInput && i == m.fc.Index {
			// Ставим фокус на элемент
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		} else {
			// Убираем фокус с элемента
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = noStyle
			m.inputs[i].TextStyle = noStyle
		}

	}

	for i := 0; i <= len(m.controlInputs)-1; i++ {
		if m.fc.Group == focusGroupControl && i == m.fc.Index {
			// Ставим фокус на элемент
			m.controlInputs[i].Focus()
			continue
		}
		m.controlInputs[i].Blur()
	}

	return tea.Batch(cmds...)
}

func (m *LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	// logger.Get().Println("updates")
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	for i := range m.controlInputs {
		_, cmds[i] = m.controlInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *LoginModel) isValid() bool {
	// login := ""
	// password := ""
	// repeatedPassword := ""

	// for _, i := range m.inputs {
	// 	switch i.Type {
	// 	case inputLogin:
	// 		login = i.Value()
	// 	case inputPassword:
	// 		password = i.Value()
	// 	case inputRepeatedPassword:
	// 		repeatedPassword = i.Value()
	// 	}
	// }

	login := m.getValueFromInput(inputLogin)
	if login == "" {
		m.errorMessage.Show = true
		m.errorMessage.Message = "пустой логин"
		return false
	}

	password := m.getValueFromInput(inputPassword)
	if password == "" {
		m.errorMessage.Show = true
		m.errorMessage.Message = "пустой пароль"
		return false
	}

	repeatedPassword := m.getValueFromInput(inputRepeatedPassword)
	if m.registration && password != repeatedPassword {
		m.errorMessage.Show = true
		m.errorMessage.Message = "пароли не совпадают"
		return false
	}

	return true
}

func (m *LoginModel) getValueFromInput(t int) string {
	for i, input := range m.inputs {
		if input.Type == t {
			return m.inputs[i].Value()
		}
	}

	return ""
}

func NewLoginTextInput() TypedTextInput {
	item := TypedTextInput{
		Model: textinput.New(),
		Type:  inputLogin,
	}
	item.Cursor.SetMode(cursor.CursorBlink)
	item.CharLimit = 32
	item.Placeholder = "логин"
	item.TextStyle = focusedStyle
	item.PromptStyle = focusedStyle
	item.Focus()
	return item
}

func NewPasswordTextInput() TypedTextInput {
	item := TypedTextInput{
		Model: textinput.New(),
		Type:  inputPassword,
	}
	item.Cursor.SetMode(cursor.CursorBlink)
	item.CharLimit = 32
	item.Placeholder = "пароль"
	return item
}

func NewRepeatPasswordTextInput() TypedTextInput {
	item := TypedTextInput{
		Model: textinput.New(),
		Type:  inputRepeatedPassword,
	}
	item.Cursor.SetMode(cursor.CursorBlink)
	item.CharLimit = 32
	item.Placeholder = "повторите пароль"
	return item
}

func (m *LoginModel) login() (tea.Model, tea.Cmd) {
	err := core.Instance.Login(m.getValueFromInput(inputLogin), m.getValueFromInput(inputPassword))
	if err != nil {
		logger.Get().Printf("login error: %s\n", err.Error())
		m.errorMessage.Show = true
		m.errorMessage.Message = "internal error"
		return m, nil
	}
	logger.Get().Println(core.Instance.User)
	return m, func() tea.Msg { return messages.LoginSuccessMsg{} }

}
