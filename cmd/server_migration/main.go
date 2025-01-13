package main

import "github.com/besean163/gophkeeper/internal/server/database"

func main() {
	err := database.RunMigrations("postgres://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable")
	if err != nil {
		panic(err)
	}
}
