package storage

import (
	"database/sql"
	"github.com/GrayC9/TaskManager/internal/handlers"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {

		log.Printf("Ошибка подключения к базе данных: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("Не удалось подключиться к базе данных: %v", err)
	}

	log.Printf("Успешное подключение к базе данных!")
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SaveTaskToDB(task handlers.TaskData) error {
	query := "INSERT INTO tasks (title, description, severity_id, employee_id) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, task.Title, task.Description, task.SeverityID, task.EmployeeID)
	if err != nil {
		log.Printf("Ошибка сохранения задачи в базу данных: %v", err)
	}
	return nil
}
