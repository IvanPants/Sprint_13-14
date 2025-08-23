package main

import (
	"log"
	"os"

	"sprint_13-14/pkg/db"
)

func main() {

	dbFile := "scheduler.db"
	if err := db.Init(dbFile); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	port := getPort()

	if err := StartServer(port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func getPort() string {

	if port := os.Getenv("TODO_PORT"); port != "" {
		return port
	}

	return "7540"
}
