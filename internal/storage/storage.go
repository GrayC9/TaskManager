package storage

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitDB(dsn string) error {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SaveToDB(lastName, firstName, threeName, tekstObrasheniya, timeToComplete string) error {
	query := "INSERT INTO tasks (last_name, first_name, three_name, tekst_obrasheniya, time_to_complete) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, lastName, firstName, threeName, tekstObrasheniya, timeToComplete)
	if err != nil {
		return fmt.Errorf("Ошибка сохранения в базу данных: %v", err)
	}
	return nil
}
