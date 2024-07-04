package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateComment(comment *Comment) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO comments (comment_id, user_id, content, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(comment.PostID, comment.UserID, comment.Content, time.Now())
	return err
}

func UpdateCommentLikes(commentID int, likes int, dislikes int) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE comments SET likes = ?, dislikes = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(likes, dislikes, commentID)
	return err
}

func GetCommentByID(commentID int) (*Comment, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var comment Comment
	err = db.QueryRow("SELECT id, post_id, user_id, content, likes, dislikes, created_at FROM comments WHERE id = ?", commentID).Scan(
		&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes, &comment.Dislikes, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
