package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer(port string) error {
	webDir := "./web"

	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		return fmt.Errorf("директория %s не существует", webDir)
	}

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s", port)
	log.Printf("Откройте http://localhost:%s в браузере", port)

	return http.ListenAndServe(":"+port, nil)
}
