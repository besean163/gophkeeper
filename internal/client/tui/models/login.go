package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
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

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Render("[ Вход ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Вход"))

	errorStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).BorderForeground(lipgloss.Color("#FF5A40"))
)

type LoginModel struct {
	registration bool
	focusIndex   int
	showError    bool
	inputs       []textinput.Model
	errorMessage string
}

func NewLoginModel(registration bool) *LoginModel {
	inputCount := 2
	if registration {
		inputCount = 3
	}

	inputs := make([]textinput.Model, inputCount)
	var input textinput.Model
	for i := range inputs {
		input = textinput.New()
		input.Cursor.SetMode(cursor.CursorBlink)
		input.CharLimit = 32

		switch i {
		case inputLogin:
			input.Placeholder = "login"
			input.TextStyle = focusedStyle
			input.PromptStyle = focusedStyle
			input.Focus()
		case inputPassword:
			input.Placeholder = "password"
		case inputRepeatedPassword:
			input.Placeholder = "repeat password"
		}
		inputs[i] = input
	}

	return &LoginModel{
		registration: registration,
		inputs:       inputs,
		showError:    true,
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))
	// logger.Get().Println("login update")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// смена фокуса
		case "tab", "shift+tab", "enter", "up", "down":
			cmds = append(cmds, m.moveFocus(msg))
		}
	}

	cmds = append(cmds, m.updateError())

	// Handle character input and blinking
	cmds = append(cmds, m.updateInputs(msg))

	return m, tea.Batch(cmds...)
}

func (m *LoginModel) View() string {
	// logger.Get().Println("login view")

	var b strings.Builder
	b.WriteRune('\n')
	b.WriteString("Введите ваши данные:\n")

	// logger.Get().Println(len(m.inputs))
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	if m.showError {
		logger.Get().Println("here")
		b.WriteRune('\n')
		b.WriteString(errorStyle.Render(m.errorMessage))
		b.WriteRune('\n')
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m *LoginModel) moveFocus(msg tea.KeyMsg) tea.Cmd {
	s := msg.String()

	if s == "enter" && m.focusIndex == len(m.inputs) {
		logger.Get().Println("submit")
		return m.Submit()
	}
	logger.Get().Println("here")

	if s == "up" || s == "shift+tab" {
		m.focusIndex--
	} else {
		m.focusIndex++
	}

	if m.focusIndex > len(m.inputs) {
		// оставляем на кнопке submit
		m.focusIndex = len(m.inputs)
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.inputs)
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			// Ставим фокус на элемент
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		}
		// Убираем фокус с элемента
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}

	return tea.Batch(cmds...)
}

func (m *LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	// logger.Get().Println("updates")
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *LoginModel) updateError() tea.Cmd {
	if !m.showError {
		return nil

	}
	return func() tea.Msg {
		time.Sleep(time.Second * 1)
		m.showError = false
		return struct{}{}
	}
}

func (m *LoginModel) Submit() tea.Cmd {

	login := m.inputs[inputLogin].Value()
	password := m.inputs[inputPassword].Value()
	repeatPassword := m.inputs[inputRepeatedPassword].Value()

	if login == "" {
		m.errorMessage = "Пустой логин."
		m.showError = true
	} else if password == "" {
		m.errorMessage = "Пустой пароль."
		m.showError = true
	} else if password != repeatPassword {
		m.errorMessage = "Пароли не совпадают."
		m.showError = true
	} else {
		d := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{
			Login:    login,
			Password: password,
		}
		body, _ := json.Marshal(d)

		r, err := http.NewRequest(http.MethodPost, "localhost:8080/register", body)
		m.errorMessage = "Пароли не совпадают."
		m.showError = true
	}
	logger.Get().Println(m.inputs[inputLogin].Value())
	return func() tea.Msg { return struct{}{} }
}
