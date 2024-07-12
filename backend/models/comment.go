package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
	// Likes     int       `json:"likes"`
	// Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
	Score     int       `json:"score"`
}

// CreateComment creates a new comment in the database.
func CreateComment(db *sql.DB, comment *Comment) error {
	query := `INSERT INTO comments (post_id, user_id, context, score, created_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)`
	result, err := db.Exec(query, comment.PostID, comment.UserID, comment.Content, comment.Score)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	comment.ID = int(id)
	return nil
}

// GetCommentByID retrieves a comment by its ID.
func GetCommentByID(db *sql.DB, id int) (*Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at, score FROM comments WHERE id = ?`
	row := db.QueryRow(query, id)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.Score)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentsByPostID retrieves all comments for a given post ID.
func GetCommentsByPostID(db *sql.DB, postID int) ([]Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at, score FROM comments WHERE post_id = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.Score)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// UpdateComment updates an existing comment in the database.
func UpdateCommentScore(db *sql.DB, commentID int, score int) error {
	query := `UPDATE comments SET score = ? WHERE id = ?`
	_, err := db.Exec(query, score, commentID)
	return err
}

// DeleteComment deletes a comment from the database.
func DeleteComment(db *sql.DB, id int) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
