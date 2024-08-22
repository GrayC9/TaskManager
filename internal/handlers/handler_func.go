package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrayC9/TaskManager/internal/storage"
)

type FormData struct {
	LastName         string `json:"last_name"`
	FirstName        string `json:"first_name"`
	ThreeName        string `json:"three_name"`
	TekstObrasheniya string `json:"tekst_obrasheniya"`
	TimeToComplete   string `json:"time_to_complete"`
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы поддерживаются", http.StatusMethodNotAllowed)
		return
	}

	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Получены данные: %+v\n", data)

	err = storage.SaveToDB(data.LastName, data.FirstName, data.ThreeName, data.TekstObrasheniya, data.TimeToComplete)
	if err != nil {
		http.Error(w, "Не удалось сохранить данные в базу данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные успешно сохранены"))
}
