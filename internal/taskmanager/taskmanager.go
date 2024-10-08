package taskmanager

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GrayC9/TaskManager/internal/config"
	"github.com/GrayC9/TaskManager/internal/handlers"
	"github.com/GrayC9/TaskManager/internal/storage"
)

func Run() {
	cfg := config.LoadConfig()

	err := storage.InitDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer storage.CloseDB()

	http.HandleFunc("/", handlers.TaskHandler)

	fmt.Println("Сервер запущен на порту", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
