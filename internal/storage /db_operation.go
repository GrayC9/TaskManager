package storage

import (
	"Gazprome_TaskManager/internal/models"
	"database/sql"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := "INSERT INTO tasks (requestor, description, deadline, status, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, task.Requestor, task.Description, task.Deadline, task.Status, task.CreatedAt)
	return err
}

func (r *TaskRepository) GetTasks() ([]*models.Task, error) {
	rows, err := r.DB.Query("SELECT id, requestor, description, deadline, status, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Requestor, &task.Description, &task.Deadline, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTaskStatus(id int, status string) error {
	query := "UPDATE tasks SET status = ? WHERE id = ?"
	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
