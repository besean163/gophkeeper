package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
}

func (m LoginModel) Init() tea.Cmd {
	return nil
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println("login", msg)

	return m, nil
}

func (m LoginModel) View() string {
	return ""
}
