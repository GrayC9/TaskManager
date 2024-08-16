package main

import (
	"log"
	"net/http"

	"Gazprome_TaskManager/internal/handlers"
	"Gazprome_TaskManager/internal/services"
	"Gazprome_TaskManager/internal/storage"
	"Gazprome_TaskManager/pkg"

	"github.com/gorilla/mux"
)

func main() {
	database, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	taskRepo := &storage.TaskRepository{DB: database}
	taskService := &services.TaskService{Repo: taskRepo}
	taskHandler := &handlers.TaskHandler{Service: taskService}

	r := mux.NewRouter()

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTaskStatus).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
