package models

import (
	"reflect"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/core"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type State int

const (
	rootSignState State = iota
	rootLoginState
	rootRegistrationState
	rootSectionsState
	rootAccountListState
	rootAccountEditState
)

type RootModel struct {
	Core core.Core
	State
	*SignModel
	*LoginModel
	*RegistrationModel
	*SectionListModel
	*AccountListModel
	*AccountEditModel
	Quit bool
}

func NewRootModel() RootModel {
	return RootModel{
		State:     rootSignState,
		SignModel: NewSignModel(),
		// LoginModel:    NewLoginModel(true),
		// SectionsModel: NewSectionModel(),
		// AccountListModel: NewAccountListModel(),
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// logger.Get().Println("root update")
	logger.Get().Println(reflect.TypeOf(msg))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			// logger.Get().Println("quit")
			m.Quit = true
			return m, tea.Quit
		}
	case messages.SignLoginMsg:
		m.clearModels()
		m.State = rootLoginState
		m.LoginModel = NewLoginModel()
	case messages.SignRegistrationMsg:
		m.clearModels()
		m.State = rootRegistrationState
		m.RegistrationModel = NewRegistrationModel()
	case messages.SignBackMsg:
		m.clearModels()
		m.State = rootSignState
		m.SignModel = NewSignModel()
	case messages.LoginSuccessMsg:
		m.clearModels()
		m.State = rootSectionsState
		m.SectionListModel = NewSectionListModel()
	case messages.SectionBackMsg:
		m.clearModels()
		m.State = rootSectionsState
		m.SectionListModel = NewSectionListModel()
	case messages.SelectAccountMsg:
		m.clearModels()
		m.State = rootAccountListState
		m.AccountListModel = NewAccountListModel()
	case messages.AccountListBackMsg:
		m.clearModels()
		m.State = rootAccountListState
		m.AccountListModel = NewAccountListModel()
	case messages.AccountEditMsg:
		m.clearModels()
		m.State = rootAccountEditState
		m.AccountEditModel = NewAccountEditModel(msg.Account)
	}

	var cmd tea.Cmd
	switch m.State {
	case rootSignState:
		_, cmd = m.SignModel.Update(msg)
	case rootLoginState:
		_, cmd = m.LoginModel.Update(msg)
	case rootRegistrationState:
		_, cmd = m.RegistrationModel.Update(msg)
	case rootSectionsState:
		_, cmd = m.SectionListModel.Update(msg)
	case rootAccountListState:
		_, cmd = m.AccountListModel.Update(msg)
	case rootAccountEditState:
		_, cmd = m.AccountEditModel.Update(msg)
	}

	return m, cmd
}

func (m RootModel) View() string {
	if m.Quit {
		return "Bye!"
	}
	result := strings.Builder{}
	switch m.State {
	case rootSignState:
		result.WriteString(m.SignModel.View())
	case rootLoginState:
		result.WriteString(m.LoginModel.View())
	case rootRegistrationState:
		result.WriteString(m.RegistrationModel.View())
	case rootSectionsState:
		result.WriteString(m.SectionListModel.View())
	case rootAccountListState:
		result.WriteString(m.AccountListModel.View())
	case rootAccountEditState:
		result.WriteString(m.AccountEditModel.View())
	}

	result.WriteRune('\n')
	result.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("ctrl+c: exit\n"))
	return result.String()
}

func (m *RootModel) clearModels() {
	m.SignModel = nil
	m.LoginModel = nil
	m.RegistrationModel = nil
	m.SectionListModel = nil
	m.AccountListModel = nil
	m.AccountEditModel = nil
}
