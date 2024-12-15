package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RootModel struct {
	SignModel
	LoginModel
	SectionsModel
	AccountsModel
}

func NewRootModel() RootModel {
	return RootModel{
		SignModel:     SignModel{},
		LoginModel:    LoginModel{},
		SectionsModel: SectionsModel{},
		AccountsModel: AccountsModel{},
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println("root", msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	}
	m.LoginModel.Update(msg)
	return m, nil
}

func (m RootModel) View() string {
	result := "\n"
	result += lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+c: exit\n")
	return result
}
