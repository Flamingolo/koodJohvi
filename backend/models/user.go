package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string    `json:"-"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"-"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// Error handling needed

func CreateUser(db *sql.DB, u User) error {
	db, err := sql.Open("sqlite3", "forum.db")

	u.ID = uuid.New().String()

	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(email, username, password_hash) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.Email, u.Username, u.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}
