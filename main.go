package main

import (
	"log"
	"os"
)

func main() {
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
