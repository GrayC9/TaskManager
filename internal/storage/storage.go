package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SaveTaskToDB(title, description string, severityID, employeeID int) error {
	query := "INSERT INTO tasks (title, description, severity_id, employee_id) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, title, description, severityID, employeeID)
	if err != nil {
		fmt.Printf("Ошибка при сохранении задачи: %v\n", err)
		return fmt.Errorf("Ошибка сохранения задачи в базу данных: %v", err)
	}
	return nil
}
