package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "backend/database/database.sql")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
