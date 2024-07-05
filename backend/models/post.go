package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePost(post *Post) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts (user_id, title, content, created_at) VALUE (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(post.UserID, post.Title, post.Content, time.Now())
	return err
}

func UpdatePostLikes(postID int, likes int, dislikes int) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE posts SET likes = ?, dislikes = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(likes, dislikes, postID)
	return err
}

func GetPostByID(postID int) (*Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var post Post
	err = db.QueryRow("SELECT id, user_id, title, content, likes, dislikes, created_at FROM posts WHERE id = ?", postID).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetAllPosts () ([]Post, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, user_id, title, content, likes, dislikes, (likes - dislikes) as score, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next(){
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.Score, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}