package models

import (
	"reflect"
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
	LoginButtonBack
)

type LoginModel struct {
	fc           *components.GroupFocusCursor
	inputs       []components.TextInputModel
	buttons      []components.ButtonModel
	errorMessage *components.ErrorMessageModel
}

type ButtonEnterMsg struct{}
type ButtonBackMsg struct{}

func NewLoginModel() *LoginModel {
	item := &LoginModel{
		fc:           components.NewGroupFocusCursor(LoginGroupInput, 0),
		inputs:       getInputs(),
		buttons:      getButtons(),
		errorMessage: &components.ErrorMessageModel{},
	}
	item.activateButtons()
	return item
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("login update")
	logger.Get().Println(reflect.TypeOf(msg))
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
			logger.Get().Println("update")
			m.moveFocus(msg)
		}
	case ButtonEnterMsg:
		logger.Get().Println("enter msg")
	case ButtonBackMsg:
		logger.Get().Println("back msg")
	}
	m.updateInputs(msg)

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
		b.WriteString(button.View())
		b.WriteRune('\n')
	}
	return b.String()
}

func (m *LoginModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

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
		if m.fc.Group == LoginGroupInput && m.fc.Index >= len(m.inputs) {
			m.fc.Move(LoginGroupButtons, 0)
		}

		if m.fc.Group == LoginGroupButtons && (m.fc.Index >= len(m.buttons) || m.fc.Index < 0) {
			m.fc.Move(LoginGroupInput, 0)
		}

		if m.fc.Group == LoginGroupInput {
			if m.fc.Index < 0 {
				m.fc.Index = 0
			}
			return m.inputs[m.fc.Index].Focus()
		}

		if m.fc.Group == LoginGroupButtons {
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

	return tea.Batch(cmds...)
}

func (m *LoginModel) activateButtons() {
	m.buttons[LoginButtonEnter].Activate(func() bool { return m.isValid() })
	m.buttons[LoginButtonBack].Activate(func() bool { return true })
}

func (m *LoginModel) pressButtons(msg tea.Msg) tea.Cmd {
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

func (m *LoginModel) isValid() bool {

	login := m.inputs[LoginInputLogin].Value()
	password := m.inputs[LoginInputPassword].Value()

	return login != "" && password != ""
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
	logger.Get().Println("view")

	var button components.ButtonModel
	items := make([]components.ButtonModel, 0)

	// добавляем кнопку вход
	button = components.NewButtonModel("Вход")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorAzure))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(ButtonEnterMsg{})
	items = append(items, button)

	// добавляем кнопку назад
	button = components.NewButtonModel("Назад")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorOrange))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(ButtonBackMsg{})
	items = append(items, button)

	return items
}