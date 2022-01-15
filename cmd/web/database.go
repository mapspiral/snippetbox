package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func setupDatabase(connectionString string) (*sql.DB, error) {
	db, errorInfo := sql.Open("mysql", connectionString)

	if errorInfo != nil {
		return nil, errorInfo
	}

	if errorInfo = db.Ping(); errorInfo != nil {
		return nil, errorInfo
	}

	return db, nil
}
