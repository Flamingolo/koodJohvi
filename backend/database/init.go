package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	// Read .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Database SQL file path
	sqlFilePath := os.Getenv("DATABASE_SQL_PATH")
	if sqlFilePath == "" {
		log.Fatal("DATABASE_SQL_PATH is not set")
	}

	// Database file path
	dbFilePath := os.Getenv("DATABASE_FILE_PATH")
	if dbFilePath == "" {
		log.Fatal("DATABASE_FILE_PATH is not set")
	}

	log.Printf("Using DATABASE_SQL_PATH: %s\n", sqlFilePath)
	log.Printf("Using DATABASE_FILE_PATH: %s\n", dbFilePath)

	// Opening db
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Open the SQL file
	file, err := os.Open(sqlFilePath)
	if err != nil {
		log.Fatal("Error opening SQL file:", err)
	}
	defer file.Close()

	// Read the file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Error getting file stats:", err)
	}
	size := stat.Size()
	content := make([]byte, size)
	_, err = file.Read(content)
	if err != nil {
		log.Fatal("Error reading file content:", err)
	}

	// Split the SQL statement and execute them
	commands := strings.Split(string(content), ";")
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command != "" {
			_, err := db.Exec(command)
			if err != nil {
				log.Fatalf("Error executing command: %s\nError: %v", command, err)
			} else {
				log.Printf("Successfully executed command: %s\n", command)
			}
		}
	}

	log.Println("Database initialized successfully")
	return db
}
