package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Context   string    `json:"context"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateComment creates a new comment in the database.
func CreateComment(db *sql.DB, comment *Comment) error {
	query := `INSERT INTO comments (post_id, user_id, context, likes, dislikes, score, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, comment.PostID, comment.UserID, comment.Context, comment.Likes, comment.Dislikes, comment.Score, time.Now())
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
	query := `SELECT id, post_id, user_id, context, likes, dislikes, score, created_at FROM comments WHERE id = ?`
	row := db.QueryRow(query, id)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Context, &comment.Likes, &comment.Dislikes, &comment.Score, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentsByPostID retrieves all comments for a given post ID.
func GetCommentsByPostID(db *sql.DB, postID int) ([]Comment, error) {
	query := `SELECT id, post_id, user_id, context, likes, dislikes, score, created_at FROM comments WHERE post_id = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Context, &comment.Likes, &comment.Dislikes, &comment.Score, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// UpdateComment updates an existing comment in the database.
func UpdateComment(db *sql.DB, comment *Comment) error {
	query := `UPDATE comments SET context = ?, likes = ?, dislikes = ? WHERE id = ?`
	_, err := db.Exec(query, comment.Context, comment.Likes, comment.Dislikes, comment.ID)
	return err
}

// DeleteComment deletes a comment from the database.
func DeleteComment(db *sql.DB, id int) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
