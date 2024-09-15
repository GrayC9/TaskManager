package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrayC9/TaskManager/internal/models"
	"github.com/GrayC9/TaskManager/internal/storage"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы поддерживаются", http.StatusMethodNotAllowed)
		return
	}

	var data models.TaskData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Получены данные: %+v\n", data)

	err = storage.SaveTaskToDB(data)
	if err != nil {
		http.Error(w, "К сожалению, не удалось создать задачу", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача успешно сохранена"))
}
