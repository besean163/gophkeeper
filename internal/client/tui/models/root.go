package models

import (
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type State int

const (
	rootSignState State = iota
	rootLoginState
	rootSectionsState
	rootAccountsState
)

type RootModel struct {
	State
	*SignModel
	LoginModel
	SectionsModel
	AccountsModel
}

func NewRootModel() RootModel {
	return RootModel{
		State:         rootSignState,
		SignModel:     NewSignModel(),
		LoginModel:    LoginModel{},
		SectionsModel: SectionsModel{},
		AccountsModel: AccountsModel{},
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("root update")

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	}

	switch m.State {
	case rootSignState:
		m.SignModel.Update(msg)
	}

	// m.LoginModel.Update(msg)
	return m, nil
}

func (m RootModel) View() string {
	logger.Get().Println("root view")
	switch m.State {
	case rootSignState:
		return m.SignModel.View()
	}
	result := "\n"
	result += lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+c: exit\n")
	return result
}
