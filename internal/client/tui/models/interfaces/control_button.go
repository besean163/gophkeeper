package interfaces

import tea "github.com/charmbracelet/bubbletea"

type ControlButton interface {
	tea.Model
	Focus()
	Blur()
}
