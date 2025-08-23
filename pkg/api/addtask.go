package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"sprint_13-14/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSONError(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSONError(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	now := time.Now()
	if err := processTaskDate(&task, now); err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"id": id}
	writeJSON(w, response, http.StatusOK)
}

func processTaskDate(task *db.Task, now time.Time) error {
	today := now.Format(DateFormat)

	if strings.ToLower(task.Date) == "today" {
		task.Date = today
	}

	if task.Date == "" {
		task.Date = today
	}

	parsedDate, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("некорректный формат даты: %s", task.Date)
	}

	if parsedDate.Before(now.Truncate(24 * time.Hour)) {

		if task.Repeat != "" {
			nextDate, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = nextDate
		} else {

			task.Date = today
		}
	}

	return nil
}

func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
	}
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{"error": message}
	writeJSON(w, response, statusCode)
}
