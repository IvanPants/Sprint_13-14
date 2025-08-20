package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StartServer запускает веб-сервер на указанном порту
func StartServer(port string) error {
	webDir := "./web"

	// Проверяем существование директории с веб-файлами
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		return fmt.Errorf("директория %s не существует", webDir)
	}

	// Настраиваем обработчик для статических файлов
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	// Запускаем сервер
	log.Printf("Сервер запущен на порту %s", port)
	log.Printf("Откройте http://localhost:%s в браузере", port)

	return http.ListenAndServe(":"+port, nil)
}
