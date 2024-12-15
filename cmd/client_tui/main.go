package main

import (
	"fmt"
	"log"

	app "github.com/besean163/gophkeeper/internal/client/tui"
)

func main() {
	// Инициализация программы
	p := app.NewProgram()

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
