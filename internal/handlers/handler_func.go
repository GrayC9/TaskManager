package handlers

import (
	"encoding/json"
	"fmt"
	"main/internal/storage"
	"net/http"
	"time"
)

type FormData struct {
	FIO              string `json:"fio"`
	TekstObrasheniya string `json:"tekst_obrasheniya"`
	TimeToComplete   string `json:"time_to_complete"`
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы ", http.StatusMethodNotAllowed)
		return
	}

	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Получены данные: %+v\n", data)

	record := fmt.Sprintf("ФИО: %s\nТекст обращения: %s\nВремя на выполнение: %s\nДата обращения: %s\n\n",
		data.FIO, data.TekstObrasheniya, data.TimeToComplete, time.Now().Format(time.RFC1123))

	err = storage.SaveToFile(record)
	if err != nil {
		http.Error(w, "Не удалось сохранить данные", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные успешно сохранены"))
}
