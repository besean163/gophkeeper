package models

import (
	"fmt"
	"strings"

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
	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type LoginModel struct {
	registration bool
	focusIndex   int
	inputs       []textinput.Model
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
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// logger.Get().Println("login update")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// смена фокуса
		case "tab", "shift+tab", "enter", "up", "down":

			s := msg.String()
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
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
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m *LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	// logger.Get().Println("updates")
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
