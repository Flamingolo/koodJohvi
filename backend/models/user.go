package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	PostCount    int       `json:"post_count"`
	CommentCount int       `json:"comment_count"`
}

type Session struct {
	UserID       int       `json:"user_id"`
	SessionId    string    `json:"session_id"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
	LastActivity time.Time `json:"last_activity"`
}

func CreateUser(db *sql.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (nickname, age, gender, first_name, last_name, email, password, created_at, post_count, comment_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, hashedPassword, time.Now(), user.PostCount, user.CommentCount)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, nickname, age, gender, first_name, last_name, email, created_at, post_count, comment_count FROM users WHERE id = ?`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.PostCount, &user.CommentCount)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AuthenticateUser(db *sql.DB, identifier, password string) (*User, error) {
	query := `SELECT id, nickname, age, gender, first_name, last_name, email, password, created_at, post_count, comment_count FROM users WHERE nickname = ? OR email = ?`
	row := db.QueryRow(query, identifier, identifier)

	var user User
	var hashedPassword string
	err := row.Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &hashedPassword, &user.CreatedAt, &user.PostCount, &user.CommentCount)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateSession(db *sql.DB, userID int) (*Session, error) {
	sessionID := uuid.New().String()
	createdAt := time.Now()
	expiresAt := createdAt.Add(24 * time.Hour)

	query := `INSERT INTO active_sessions (user_id, session_id, created_at, expires_at, last_activity) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, userID, sessionID, createdAt, expiresAt, createdAt)
	if err != nil {
		return nil, err
	}

	session := &Session{
		UserID:       userID,
		SessionId:    sessionID,
		CreatedAt:    createdAt,
		ExpiresAt:    expiresAt,
		LastActivity: createdAt,
	}

	return session, nil
}

func ValidateSession(db *sql.DB, sessionID string) (*Session, error) {
	query := `SELECT user_id, session_id, created_at, expires_at, last_activity FROM active_sessions WHERE session_id = ? AND expires_at > ?`
	row := db.QueryRow(query, sessionID, time.Now())

	var session Session
	err := row.Scan(&session.UserID, &session.SessionId, &session.CreatedAt, &session.ExpiresAt, &session.LastActivity)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func UpdateSessionActivity(db *sql.DB, sessionID string) error {
	query := `UPDATE active_sessions SET last_activity = ? WHERE session_id = ?`
	_, err := db.Exec(query, time.Now(), sessionID)
	return err
}

func DeleteSession(db *sql.DB, sessionID string) error {
	query := `DELETE FROM active_sessions WHERE session_id = ?`
	_, err := db.Exec(query, sessionID)
	return err
}
