package main

import (
	"fmt"
	"log"

	app "github.com/besean163/gophkeeper/internal/client/tui"
	"github.com/besean163/gophkeeper/internal/client/tui/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := models.NewRootModel()
	p := tea.NewProgram(m)
	// Инициализация программы
	// p := app.NewProgram()

	// Запуск программы
	finalModel, err := p.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске программы: %v", err)
	}

	// Вывод результата (например, финального состояния программы)
	if finalModel, ok := finalModel.(app.Model); ok {
		if finalModel.LoggedIn {
			fmt.Println("Успешный вход в систему.")
		} else {
			fmt.Println("Программа завершена без входа.")
		}
	}
}
