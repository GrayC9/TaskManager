package services

import (
	"Gazprome_TaskManager/internal/models"
	"Gazprome_TaskManager/internal/storage"
)

type TaskService struct {
	Repo *storage.TaskRepository
}

func (s *TaskService) CreateTask(task *models.Task) error {
	task.Status = "скоро приступлю"
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]*models.Task, error) {
	return s.Repo.GetTasks()
}

func (s *TaskService) UpdateTaskStatus(id int, status string) error {
	return s.Repo.UpdateTaskStatus(id, status)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.Repo.DeleteTask(id)
}
