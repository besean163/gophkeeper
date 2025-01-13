// Package app пакет предоставляет абстракцию приложения Tui
package app

import (
	"github.com/besean163/gophkeeper/internal/client/interfaces"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	"github.com/besean163/gophkeeper/internal/logger"
	tea "github.com/charmbracelet/bubbletea"
)

// App структура приложения
type App struct {
	rootModel tea.Model
}

// NewApp создание структуры приложения
func NewApp(core interfaces.Core, logger logger.Logger) App {
	return App{
		rootModel: models.NewRootModel(core, logger),
	}
}

// Run запуск приложения
func (app App) Run() error {
	program := tea.NewProgram(app.rootModel)

	_, err := program.Run()
	if err != nil {
		return err
	}

	return nil
}
