package models

import (
	"reflect"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	keybinding "github.com/besean163/gophkeeper/internal/client/tui/models/key_binding"
	"github.com/besean163/gophkeeper/internal/logger"
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
	rootNoteListState
	rootNoteEditState
	rootCardListState
	rootCardEditState
)

// RootModel корневая модель
type RootModel struct {
	logger logger.Logger
	core   interfaces.Core
	State
	*SignModel
	*LoginModel
	*RegistrationModel
	*SectionListModel
	*AccountListModel
	*AccountEditModel
	*NoteListModel
	*NoteEditModel
	*CardListModel
	*CardEditModel
	Quit bool
}

func NewRootModel(core interfaces.Core, logger logger.Logger) RootModel {
	return RootModel{
		logger:    logger,
		core:      core,
		State:     rootSignState,
		SignModel: NewSignModel(logger),
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.logger.Debug("msg type", logger.NewField("type", reflect.TypeOf(msg)))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == keybinding.Ctrlc {
			m.Quit = true
			return m, tea.Quit
		}
	case messages.SignLoginMsg:
		m.clearModels()
		m.State = rootLoginState
		m.LoginModel = NewLoginModel(m.core, m.logger)
	case messages.SignRegistrationMsg:
		m.clearModels()
		m.State = rootRegistrationState
		m.RegistrationModel = NewRegistrationModel(m.core, m.logger)
	case messages.SignBackMsg:
		m.clearModels()
		m.State = rootSignState
		m.SignModel = NewSignModel(m.logger)
	case messages.LoginSuccessMsg:
		m.clearModels()
		m.State = rootSectionsState
		m.SectionListModel = NewSectionListModel(m.logger)
	case messages.RegistrationSuccessMsg:
		m.clearModels()
		m.State = rootSectionsState
		m.SectionListModel = NewSectionListModel(m.logger)
	case messages.SectionBackMsg:
		m.clearModels()
		m.State = rootSectionsState
		m.SectionListModel = NewSectionListModel(m.logger)
	case messages.SelectAccountMsg:
		m.clearModels()
		m.State = rootAccountListState
		m.AccountListModel = NewAccountListModel(m.core, m.logger)
	case messages.AccountListBackMsg:
		m.clearModels()
		m.State = rootAccountListState
		m.AccountListModel = NewAccountListModel(m.core, m.logger)
	case messages.AccountEditMsg:
		m.clearModels()
		m.State = rootAccountEditState
		m.AccountEditModel = NewAccountEditModel(m.core, msg.Account, m.logger)

	case messages.SelectNoteMsg:
		m.clearModels()
		m.State = rootNoteListState
		m.NoteListModel = NewNoteListModel(m.core, m.logger)
	case messages.NoteListBackMsg:
		m.clearModels()
		m.State = rootNoteListState
		m.NoteListModel = NewNoteListModel(m.core, m.logger)
	case messages.NoteEditMsg:
		m.clearModels()
		m.State = rootNoteEditState
		m.NoteEditModel = NewNoteEditModel(m.core, msg.Note, m.logger)

	case messages.SelectCardMsg:
		m.clearModels()
		m.State = rootCardListState
		m.CardListModel = NewCardListModel(m.core, m.logger)
	case messages.CardListBackMsg:
		m.clearModels()
		m.State = rootCardListState
		m.CardListModel = NewCardListModel(m.core, m.logger)
	case messages.CardEditMsg:
		m.clearModels()
		m.State = rootCardEditState
		m.CardEditModel = NewCardEditModel(m.core, msg.Card, m.logger)
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
	case rootNoteListState:
		_, cmd = m.NoteListModel.Update(msg)
	case rootNoteEditState:
		_, cmd = m.NoteEditModel.Update(msg)
	case rootCardListState:
		_, cmd = m.CardListModel.Update(msg)
	case rootCardEditState:
		_, cmd = m.CardEditModel.Update(msg)
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
	case rootNoteListState:
		result.WriteString(m.NoteListModel.View())
	case rootNoteEditState:
		result.WriteString(m.NoteEditModel.View())
	case rootCardListState:
		result.WriteString(m.CardListModel.View())
	case rootCardEditState:
		result.WriteString(m.CardEditModel.View())
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
