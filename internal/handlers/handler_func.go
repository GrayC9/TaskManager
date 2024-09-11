package handlers

import (
	"encoding/json"
	"github.com/GrayC9/TaskManager/internal/storage"
	"log"
	"net/http"
)

type TaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SeverityID  int    `json:"severity_id"`
	EmployeeID  int    `json:"employee_id"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы поддерживаются", http.StatusMethodNotAllowed)
		return
	}

	var data TaskData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Получены данные: %+v\n", data)

	err = storage.SaveTaskToDB(data)
	if err != nil {
		http.Error(w, "К солжалению, не удалость создать задачу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача успешно сохранена"))
}
