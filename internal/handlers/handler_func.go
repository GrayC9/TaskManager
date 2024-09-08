package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrayC9/TaskManager/internal/storage"
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

	fmt.Printf("Получены данные: %+v\n", data)

	err = storage.SaveTaskToDB(data.Title, data.Description, data.SeverityID, data.EmployeeID)
	if err != nil {
		http.Error(w, "Не удалось сохранить задачу в базу данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача успешно сохранена"))
}
