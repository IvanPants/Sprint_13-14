package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"sprint_13-14/pkg/api" // Импортируем пакет api
)

// StartServer запускает веб-сервер на указанном порту
func StartServer(port string) error {
	webDir := "./web"

	// Проверяем существование директории с веб-файлами
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		return fmt.Errorf("директория %s не существует", webDir)
	}

	api.Init()

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s", port)
	log.Printf("Откройте http://localhost:%s в браузере", port)

	return http.ListenAndServe(":"+port, nil)
}
