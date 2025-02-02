package models

import (
	"strings"

	coremodels "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/messages"
	"github.com/besean163/gophkeeper/internal/client/tui/models/components"
	keybinding "github.com/besean163/gophkeeper/internal/client/tui/models/key_binding"
	"github.com/besean163/gophkeeper/internal/client/tui/models/styles"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	AccountEditGroupInput = iota
	AccountEditGroupButtons
)

const (
	AccountEditInputName = iota
	AccountEditInputLogin
	AccountEditInputPassword
)

const (
	AccountEditButtonSave = iota
	AccountEditButtonBack
)

// AccountEditModel модель окна редактирования аккаунта
type AccountEditModel struct {
	logger  logger.Logger
	core    interfaces.Core
	account coremodels.Account
	fc      *components.GroupFocusCursor
	inputs  []components.TextInputModel
	buttons []components.ButtonModel
}

func NewAccountEditModel(core interfaces.Core, account coremodels.Account, logger logger.Logger) *AccountEditModel {
	item := &AccountEditModel{
		logger:  logger,
		core:    core,
		fc:      components.NewGroupFocusCursor(AccountEditGroupInput, 0),
		account: account,
	}

	item.setInputs()
	item.setButtons()
	return item
}

func (m *AccountEditModel) Init() tea.Cmd {
	return nil
}

func (m *AccountEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.activateButtons()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Tab, keybinding.ShiftTab, keybinding.Enter, keybinding.Up, keybinding.Down:
			var cmd tea.Cmd
			cmd = m.pressButtons(msg)
			if cmd != nil {
				return m, cmd
			}
			m.logger.Info("update")
			m.moveFocus(msg)
		}
	case messages.ButtonSubmitMsg:
		m.logger.Info("save msg")
		return m, m.save()
	case messages.ButtonBackMsg:
		m.logger.Info("back msg")
		return m, func() tea.Msg { return messages.AccountListBackMsg{} }
	}
	m.updateInputs(msg)
	m.activateButtons()
	return m, nil
}

func (m *AccountEditModel) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().PaddingLeft(2).Render("Редактирование аккаунта"))
	b.WriteRune('\n')
	b.WriteRune('\n')
	for _, input := range m.inputs {
		b.WriteString(lipgloss.NewStyle().PaddingLeft(4).Render(input.View()))
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for _, button := range m.buttons {
		b.WriteString(button.View())
		b.WriteByte('\n')
	}
	return b.String()
}

func (m *AccountEditModel) moveFocus(msg tea.KeyMsg) tea.Cmd {

	for i := range m.inputs {
		m.inputs[i].Blur()
	}

	// проставляем активность на кнопках, чтобы не активные не выбирать
	m.activateButtons()
	for i := range m.buttons {
		m.buttons[i].Blur()
	}

	for {
		key := msg.String()
		if key == keybinding.Up || key == keybinding.ShiftTab {
			m.fc.Index--
		} else {
			m.fc.Index++
		}

		// переключить на другой список если требуется и установить верный индекс
		if m.fc.Group == AccountEditGroupInput && m.fc.Index >= len(m.inputs) {
			m.fc.Move(AccountEditGroupButtons, 0)
		}

		if m.fc.Group == AccountEditGroupButtons {
			if m.fc.Index >= len(m.buttons) {
				m.fc.Move(AccountEditGroupInput, 0)
			} else if m.fc.Index < 0 {
				m.fc.Move(AccountEditGroupInput, len(m.inputs)-1)
			}
		}

		if m.fc.Group == AccountEditGroupInput {
			if m.fc.Index < 0 {
				m.fc.Index = 0
			}
			return m.inputs[m.fc.Index].Focus()
		}

		if m.fc.Group == AccountEditGroupButtons {
			// если кнопка активна оставляем фокус
			if m.buttons[m.fc.Index].IsActive() {
				m.buttons[m.fc.Index].Focus()
				break
			}

		}
	}

	return nil
}

func (m *AccountEditModel) activateButtons() {
	m.buttons[AccountEditButtonSave].Activate(func() bool { return m.isValid() })
	m.buttons[AccountEditButtonBack].Activate(func() bool { return true })
}

func (m *AccountEditModel) pressButtons(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keybinding.Enter:
			var cmd tea.Cmd
			for _, button := range m.buttons {
				cmd = button.Press()
				if cmd != nil {
					return cmd
				}
			}
		}
	}
	return nil
}

func (m *AccountEditModel) isValid() bool {

	name := m.inputs[AccountEditInputName].Value()
	login := m.inputs[AccountEditInputLogin].Value()
	password := m.inputs[AccountEditInputPassword].Value()

	return name != "" && login != "" && password != ""
}

func (m *AccountEditModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *AccountEditModel) setInputs() {

	var input components.TextInputModel
	items := make([]components.TextInputModel, 3)

	input = components.NewTextInputModel()
	input.Focus()
	input.Prompt = "Название"
	input.Placeholder = "введите название ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.account.Name)
	items[AccountEditInputName] = input

	input = components.NewTextInputModel()
	input.Prompt = "   Логин"
	input.Placeholder = "введите логин ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.account.Login)
	items[AccountEditInputLogin] = input

	input = components.NewTextInputModel()
	input.Prompt = "  Пароль"
	input.Placeholder = "введите пароль ..."
	input.WithFocusStyle(lipgloss.NewStyle().Foreground(styles.ColorAzure))
	input.SetValue(m.account.Password)
	items[AccountEditInputPassword] = input

	m.inputs = items
}

func (m *AccountEditModel) setButtons() {
	var button components.ButtonModel
	items := make([]components.ButtonModel, 2)

	// добавляем кнопку вход
	button = components.NewButtonModel("Сохранить")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGreen))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonSubmitMsg{})
	items[AccountEditButtonSave] = button

	// добавляем кнопку назад
	button = components.NewButtonModel("Отмена")
	button.WithStyle(lipgloss.NewStyle().PaddingLeft(4))
	button.WithSelectedStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorRed))
	button.WithNotActiveStyle(lipgloss.NewStyle().PaddingLeft(4).Foreground(styles.ColorGrey))
	button.WithPressMsg(messages.ButtonBackMsg{})
	items[AccountEditButtonBack] = button

	m.buttons = items
}

func (m *AccountEditModel) save() tea.Cmd {
	account := m.account
	account.Name = m.inputs[AccountEditInputName].Value()
	account.Login = m.inputs[AccountEditInputLogin].Value()
	account.Password = m.inputs[AccountEditInputPassword].Value()

	err := m.core.SaveAccount(account)

	if err != nil {
		m.logger.Info("error")
	} else {
		m.logger.Info("save success continue")
	}

	return func() tea.Msg { return messages.AccountListBackMsg{} }
}
