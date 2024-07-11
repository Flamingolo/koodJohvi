package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePost(db *sql.DB, post *Post) error {
	query := `INSERT INTO posts (user_id, category, title, content, likes, dislikes, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(query, post.UserID, post.Category, post.Title, post.Content, post.Likes, post.Dislikes, time.Now())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)
	return nil
}

func GetPostByID(db *sql.DB, id int) (*Post, error) {
	query := `SELECT id, user_id, category, title, content, likes, dislikes, score, created_at FROM posts WHERE id = ?`
	row := db.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Score, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPosts(db *sql.DB) ([]Post, error) {
	query := `SELECT id, user_id, category, title, content, likes, dislikes, score, created_at FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Score, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func UpdatePost(db *sql.DB, post *Post) error {
	query := `UPDATE posts SET category = ?, title = ?, content = ?, likes = ?, dislikes = ? WHERE id = ?`
	_, err := db.Exec(query, post.Category, post.Title, post.Content, post.Likes, post.Dislikes, post.ID)
	return err
}

func DeletePost(db *sql.DB, id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
