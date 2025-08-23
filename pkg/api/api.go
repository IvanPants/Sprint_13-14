package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)   // Обработка одной задачи (GET, POST, PUT)
	http.HandleFunc("/api/tasks", tasksHandler) // Обработка списка задач (GET)
	http.HandleFunc("/api/task/done", taskDoneHandler) // Добавляем обработчик для отметки выполнения
}

