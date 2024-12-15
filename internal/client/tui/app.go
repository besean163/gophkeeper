package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	StateLoginMenu State = iota
	StateCredentials
	StateSections
	StateAccounts
)

type Model struct {
	State          State
	LoginMenuIndex int
	LoggedIn       bool

	// Поля для ввода логина и пароля
	LoginInput    textinput.Model
	PasswordInput textinput.Model
	ErrorMessage  string

	// Выбор кнопки
	CredentialsMenuIndex int

	// Список разделов
	Sections        []string
	SelectedSection int

	// Данные для окна "Аккаунты"
	Accounts         []Account
	FilterInput      textinput.Model
	FilteredAccounts []Account
	SelectedAccount  int
}

type Account struct {
	Name     string
	Login    string
	Password string
}

func NewProgram() *tea.Program {
	return tea.NewProgram(InitialModel())
}

func InitialModel() Model {
	loginInput := textinput.New()
	loginInput.Placeholder = "Введите логин"
	loginInput.Focus()

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Введите пароль"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.EchoCharacter = '●'

	filterInput := textinput.New()
	filterInput.Placeholder = "Фильтр по имени"

	return Model{
		State:                StateLoginMenu,
		LoginMenuIndex:       0,
		CredentialsMenuIndex: 0,
		LoggedIn:             false,
		LoginInput:           loginInput,
		PasswordInput:        passwordInput,
		ErrorMessage:         "",
		Sections:             []string{"Аккаунты"},
		Accounts:             SampleAccounts(),
		FilterInput:          filterInput,
		FilteredAccounts:     SampleAccounts(),
		SelectedAccount:      0,
	}
}

func SampleAccounts() []Account {
	return []Account{
		{"Google", "user1@gmail.com", "pass123"},
		{"GitHub", "dev@example.com", "securepwd"},
		{"Twitter", "user123", "mypassword"},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Общая горячая клавиша для выхода
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.State {
		case StateLoginMenu:
			return m.updateLoginMenu(msg)
		case StateCredentials:
			return m.updateCredentials(msg)
		case StateSections:
			return m.updateSections(msg)
		case StateAccounts:
			return m.updateAccounts(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.State {
	case StateLoginMenu:
		return m.viewLoginMenu()
	case StateCredentials:
		return m.viewCredentials()
	case StateSections:
		return m.viewSections()
	case StateAccounts:
		return m.viewAccounts()
	default:
		return "Неизвестное состояние"
	}
}

func (m Model) viewLoginMenu() string {
	options := []string{"Sign In", "Sign Out"}
	var b strings.Builder
	b.WriteString("Выберите действие (нажмите 'q' для выхода):\n\n")
	for i, option := range options {
		cursor := " " // Пустой курсор
		if m.LoginMenuIndex == i {
			cursor = ">" // Курсор на текущем элементе
		}
		b.WriteString(fmt.Sprintf("%s %s\n", cursor, option))
	}
	return b.String()
}

func (m Model) updateLoginMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.LoginMenuIndex > 0 {
			m.LoginMenuIndex--
		}
	case "down", "j":
		if m.LoginMenuIndex < 1 {
			m.LoginMenuIndex++
		}
	case "enter":
		if m.LoginMenuIndex == 0 { // Sign In
			m.State = StateCredentials
		} else { // Sign Out
			m.LoggedIn = false
		}
	}
	return m, nil
}

func (m Model) viewCredentials() string {
	options := []string{"Поле логина", "Поле пароля", "Вход"}
	var b strings.Builder
	b.WriteString("Введите данные для входа (нажмите 'q' для выхода):\n\n")
	b.WriteString(fmt.Sprintf("Логин: %s\n\n", m.LoginInput.View()))
	b.WriteString(fmt.Sprintf("Пароль: %s\n\n", m.PasswordInput.View()))
	for i, option := range options {
		cursor := " " // Пустой курсор
		if m.CredentialsMenuIndex == i {
			cursor = ">" // Курсор на текущем элементе
		}
		b.WriteString(fmt.Sprintf("%s %s\n", cursor, option))
	}
	if m.ErrorMessage != "" {
		b.WriteString(fmt.Sprintf("\n[Ошибка]: %s", m.ErrorMessage))
	}
	return b.String()
}

func (m Model) updateCredentials(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.CredentialsMenuIndex > 0 {
			m.CredentialsMenuIndex--
		}
	case "down", "j":
		if m.CredentialsMenuIndex < 2 { // "Вход" — последний элемент
			m.CredentialsMenuIndex++
		}
	case "enter":
		switch m.CredentialsMenuIndex {
		case 0: // Поле логина
			m.LoginInput.Focus()
			m.PasswordInput.Blur()
		case 1: // Поле пароля
			m.PasswordInput.Focus()
			m.LoginInput.Blur()
		case 2: // Кнопка "Вход"
			if m.LoginInput.Value() == "admin" && m.PasswordInput.Value() == "1234" {
				m.LoggedIn = true
				m.State = StateSections
			} else {
				m.ErrorMessage = "Неверный логин или пароль!"
			}
		}
	}

	// Обновление текстовых полей
	var cmd tea.Cmd
	if m.LoginInput.Focused() {
		m.LoginInput, cmd = m.LoginInput.Update(msg)
	} else if m.PasswordInput.Focused() {
		m.PasswordInput, cmd = m.PasswordInput.Update(msg)
	}
	return m, cmd
}

func (m Model) viewSections() string {
	var b strings.Builder
	b.WriteString("Доступные разделы (нажмите 'q' для выхода):\n\n")
	for i, section := range m.Sections {
		cursor := " " // Пустой курсор
		if m.SelectedSection == i {
			cursor = ">" // Курсор на текущем элементе
		}
		b.WriteString(fmt.Sprintf("%s %s\n", cursor, section))
	}
	return b.String()
}

func (m Model) updateSections(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.SelectedSection > 0 {
			m.SelectedSection--
		}
	case "down", "j":
		if m.SelectedSection < len(m.Sections)-1 {
			m.SelectedSection++
		}
	case "enter":
		if m.SelectedSection == 0 { // Аккаунты
			m.State = StateAccounts
		}
	}
	return m, nil
}

func (m Model) viewAccounts() string {
	var b strings.Builder
	b.WriteString("Аккаунты (нажмите 'q' для выхода):\n\n")
	for i, account := range m.FilteredAccounts {
		cursor := " " // Пустой курсор
		if m.SelectedAccount == i {
			cursor = ">" // Курсор на текущем элементе
		}
		b.WriteString(fmt.Sprintf("%s %s (логин: %s)\n", cursor, account.Name, account.Login))
	}
	return b.String()
}

func (m Model) updateAccounts(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.SelectedAccount > 0 {
			m.SelectedAccount--
		}
	case "down", "j":
		if m.SelectedAccount < len(m.FilteredAccounts)-1 {
			m.SelectedAccount++
		}
	case "q":
		m.State = StateSections
	}
	return m, nil
}
