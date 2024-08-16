package pkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/taskmanager"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
		return nil, err
	}

	return db, nil
}
