package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int
	Email     string
	Username  string
	Password  string
	Nickname  string
	Age       int
	Gender    string
	FirstName string
	LastName  string
}

func CreateUser(user *User) error {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (email, username, password, nickname, age, gender, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, user.Username, user.Password, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName)
	return err
}

func GetUserByEmail(email string) (*User, error) {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT id, email, username, password, nickname, age, gender, first_name, last_name FROM users WHERE email = ?", email).Scan(
		&user.ID, &user.Email, &user.Username, &user.Password, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
