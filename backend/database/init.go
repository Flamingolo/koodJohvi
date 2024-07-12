package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// Opening db
	db, err := sql.Open("sqlite3", "backend/database/database.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Open the file
	file, err := os.Open("backend/database/database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()
	content := make([]byte, size)
	_, err = file.Read(content)
	if err != nil {
		log.Fatal(err)
	}

	// Split the sql statement and execute them
	commands := strings.Split(string(content), ";")
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command != "" {
			_, err := db.Exec(command)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return db
}
