package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(db *sql.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (nickname, age, gender, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, hashedPassword, time.Now())
	if err != nil {
		return err
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, nickname, age, gender, first_name, last_name, email, created_at FROM users WHERE id = ?`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AuthenticateUser(db *sql.DB, identifier, password string) (*User, error) {
	query := `SELECT id, nickname, age, gender, first_name, last_name, email, password, created_at FROM users WHERE nickname = ? OR email = ?`
	row := db.QueryRow(query, identifier, identifier)

	var user User
	var hashedPassword string
	err := row.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &hashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
