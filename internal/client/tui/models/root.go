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
	rootAccountListState
)

type RootModel struct {
	Core core.Core
	State
	*SignModel
	*LoginModel
	*SectionsModel
	*AccountListModel
	Quit bool
}

func NewRootModel() RootModel {
	return RootModel{
		State: rootAccountListState,
		// SignModel: NewSignModel(),
		// LoginModel:    NewLoginModel(true),
		// SectionsModel: NewSectionModel(),
		AccountListModel: NewAccountListModel(),
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
		m.State = rootSectionsState
		m.SectionsModel = NewSectionModel()
		m.LoginModel = nil
	}

	var cmd tea.Cmd
	switch m.State {
	case rootSignState:
		_, cmd = m.SignModel.Update(msg)
	case rootLoginState:
		_, cmd = m.LoginModel.Update(msg)
	case rootSectionsState:
		_, cmd = m.SectionsModel.Update(msg)
	case rootAccountListState:
		_, cmd = m.AccountListModel.Update(msg)
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
	case rootSectionsState:
		result.WriteString(m.SectionsModel.View())
	case rootAccountListState:
		result.WriteString(m.AccountListModel.View())
	}

	result.WriteRune('\n')
	result.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+c: exit\n"))
	return result.String()
}
