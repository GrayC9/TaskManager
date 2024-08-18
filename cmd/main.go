package main

import (
	"Gazprome_TaskManager/internal_1/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.FormHandler)

	fmt.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
