package models

import (
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
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
	Core core.Core
	State
	*SignModel
	*LoginModel
	SectionsModel
	AccountsModel
	Quit bool
}

func NewRootModel() RootModel {
	return RootModel{
		State: rootLoginState,
		// SignModel:     NewSignModel(),
		LoginModel:    NewLoginModel(true),
		SectionsModel: SectionsModel{},
		AccountsModel: AccountsModel{},
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// logger.Get().Println("root update")
	// logger.Get().Println(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			// logger.Get().Println("quit")
			m.Quit = true
			return m, tea.Quit
		}
	case messages.SignSuccessMsg:
		m.State = rootLoginState
		m.LoginModel = NewLoginModel(msg.WithRegistration)
		m.SignModel = nil
	case messages.LoginSuccessMsg:
		m.State = rootSignState
		m.SignModel = NewSignModel()
		m.LoginModel = nil
	}

	var cmd tea.Cmd
	switch m.State {
	case rootSignState:
		_, cmd = m.SignModel.Update(msg)
	case rootLoginState:
		_, cmd = m.LoginModel.Update(msg)
	}

	return m, cmd
}

func (m RootModel) View() string {
	if m.Quit {
		return "Bye!"
	}
	// logger.Get().Println("root view")
	result := strings.Builder{}
	switch m.State {
	case rootSignState:
		result.WriteString(m.SignModel.View())
	case rootLoginState:
		result.WriteString(m.LoginModel.View())
	}

	result.WriteRune('\n')
	result.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+c: exit\n"))
	return result.String()
}
