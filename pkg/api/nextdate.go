package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dateStr string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	date, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return "", fmt.Errorf("некорректный формат даты: %s", dateStr)
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", fmt.Errorf("некорректный формат правила")
	}

	switch parts[0] {
	case "d":
		return handleDailyRule(now, date, parts)
	case "y":
		return handleYearlyRule(now, date)
	default:
		return "", fmt.Errorf("неподдерживаемый формат правила: %s", parts[0])
	}
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	nowParam := r.FormValue("now")
	dateParam := r.FormValue("date")
	repeatParam := r.FormValue("repeat")

	var now time.Time
	if nowParam == "" {
		now = time.Now()
	} else {
		var err error
		now, err = time.Parse(DateFormat, nowParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Неверный формат параметра now: %s", nowParam), http.StatusBadRequest)
			return
		}
	}

	if dateParam == "" {
		http.Error(w, "Параметр date обязателен", http.StatusBadRequest)
		return
	}

	if repeatParam == "" {
		http.Error(w, "Параметр repeat обязателен", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(now, dateParam, repeatParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(nextDate))
}

func handleDailyRule(now, date time.Time, parts []string) (string, error) {
	if len(parts) != 2 {
		return "", fmt.Errorf("некорректный формат правила d")
	}

	days, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("некорректное количество дней: %s", parts[1])
	}

	if days < 1 || days > 400 {
		return "", fmt.Errorf("количество дней должно быть от 1 до 400")
	}

	date = date.AddDate(0, 0, days)

	for !afterNow(date, now) {
		date = date.AddDate(0, 0, days)
	}

	return date.Format(DateFormat), nil
}

func handleYearlyRule(now, date time.Time) (string, error) {

	date = date.AddDate(1, 0, 0)

	for !afterNow(date, now) {
		date = date.AddDate(1, 0, 0)
	}

	return date.Format(DateFormat), nil
}

func afterNow(date, now time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return date.After(now)
}
