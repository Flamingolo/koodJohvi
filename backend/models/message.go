package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Content    string    `json:"content"`
	IsRead     bool      `json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
}

func CreateMessage(db *sql.DB, message *Message) error {
	query := `INSERT INTO messages (sender_id, receiver_id, content, is_read, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, message.SenderID, message.ReceiverID, message.Content, message.IsRead, time.Now())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	message.ID = int(id)
	message.CreatedAt = time.Now()
	return nil
}

func GetMessageByID(db *sql.DB, id int) (*Message, error) {
	query := `SELECT id, sender_id, receiver_id, content, is_read, created_at FROM messages WHERE id = ?`
	row := db.QueryRow(query, id)

	var message Message
	err := row.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.IsRead, &message.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func GetMessagesByUserID(db *sql.DB, userID int) ([]Message, error) {
	query := `SELECT id, sender_id, receiver_id, content, is_read, created_at FROM messages WHERE sender_id = ? OR receiver_id = ?`
	rows, err := db.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content, &message.IsRead, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func MarkMessageAsRead(db *sql.DB, id int) error {
	query := `UPDATE messages SET is_read = true WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
