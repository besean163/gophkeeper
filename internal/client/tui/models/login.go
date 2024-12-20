package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	LoginGroupInput = iota
	LoginGroupButtons
)

const (
	LoginInputLogin = iota
	LoginInputPassword
)

const (
	LoginButtonEnter = iota
)

type LoginModel struct {
	fc           *components.GroupFocusCursor
	inputs       []components.TextInputModel
	buttons      []components.ButtonModel
	registration bool
	errorMessage *components.ErrorMessageModel
}

func NewLoginModel(registration bool) *LoginModel {

	return &LoginModel{
		fc:           components.NewGroupFocusCursor(LoginGroupInput, 0),
		registration: registration,
		inputs:       getInputs(),
		buttons:      getButtons(),
		errorMessage: &components.ErrorMessageModel{},
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.updateInputs(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			logger.Get().Println("update")
			m.moveFocus(msg)
		}
	}

	// var cmds []tea.Cmd
	// cmds = append(cmds, m.updateInputs(msg))

	// return m, tea.Batch(cmds...)
	return m, nil
}

func (m *LoginModel) View() string {
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
		b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(button.View()))
		b.WriteRune('\n')
	}
	return b.String()
}

func (m *LoginModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

	for i := range m.inputs {
		m.inputs[i].Blur()
	}

	for i := range m.inputs {
		m.buttons[i].Blur()
	}

	key := msg.String()
	if key == "up" || key == "shift+tab" {
		m.fc.Index--
	} else {
		m.fc.Index++
	}

	// переключить на другой список если требуется и установить верный индекс
	if m.fc.Group == LoginGroupInput && m.fc.Index > len(m.inputs) {
		m.fc.Move(LoginGroupButtons, 0)
	}

	if m.fc.Group == LoginGroupButtons && m.fc.Index < 0 {
		m.fc.Move(LoginGroupInput, len(m.inputs))
	}

	if m.fc.Group == LoginGroupInput {
		if m.fc.Index < 0 {
			m.fc.Index = 0
		}
		return m.inputs[m.fc.Index].Focus()
	}

	if m.fc.Group == LoginGroupButtons {
		for {
			// сбрасываем на инпуты
			if m.fc.Index >= len(m.buttons) {
				m.fc.Group = LoginGroupInput
				m.fc.Index = len(m.inputs) - 1
				break
			}
			// если кнопка активна оставляем фокус
			if m.buttons[m.fc.Index].IsActive() {
				logger.Get().Println("here")
				m.buttons[m.fc.Index].Focus()
				break
			}
		}
	}

	return nil

}

func (m *LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	// logger.Get().Println("updates")
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	m.buttons[LoginButtonEnter].Activate(func() bool { return m.isValid() })

	return tea.Batch(cmds...)
}

func (m *LoginModel) isValid() bool {

	login := m.inputs[LoginInputLogin].Value()
	if login == "" {
		return false
	}

	password := m.inputs[LoginInputPassword].Value()
	if password == "" {
		return false
	}

	return true
}

func getInputs() []components.TextInputModel {

	var input components.TextInputModel
	items := make([]components.TextInputModel, 0)
	input = components.NewTextInputModel()
	input.Focus()
	input.Prompt = " Логин"
	input.Placeholder = "введите логин ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	items = append(items, input)

	input = components.NewTextInputModel()
	input.Prompt = "Пароль"
	input.Placeholder = "введите пароль ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	items = append(items, input)

	return items
}

func getButtons() []components.ButtonModel {

	var button components.ButtonModel
	items := make([]components.ButtonModel, 0)
	button = components.NewButtonModel("Вход")
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	items = append(items, button)

	return items
}
