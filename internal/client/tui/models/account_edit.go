package models

import (
	"strings"

	coremodels "github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	"github.com/besean163/gophkeeper/internal/client/tui/models/interfaces"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountEditModel struct {
	account       coremodels.Account
	controlInputs []interfaces.ControlButton
	edit          bool
}

func NewAccountEditModel() *AccountEditModel {
	item := &AccountEditModel{}

	controlInputs := make([]interfaces.ControlButton, 0)
	controlInputs = append(controlInputs, components.NewBackButtonModel("Назад"))
	item.controlInputs = controlInputs
	return item
}

func (m *AccountEditModel) Init() tea.Cmd {
	logger.Get().Println("init")
	return nil
}

func (m *AccountEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.Get().Println("update")
	return m, nil
}

func (m *AccountEditModel) View() string {
	var b strings.Builder
	b.WriteString("Название: ")
	b.WriteString(m.account.Name)
	b.WriteByte('\n')

	b.WriteString("Логин: ")
	b.WriteString(m.account.Login)
	b.WriteByte('\n')

	b.WriteString("Пароль: ")
	b.WriteString(m.account.Name)
	b.WriteByte('\n')

	for _, button := range m.controlInputs {
		b.WriteString(button.View())
	}
	return b.String()
}
