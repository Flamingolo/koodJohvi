package models

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	CreatedAt        time.Time `json:"created_at"`
	AmountOfComments int       `json:"amount_of_comments"`
	Score            int       `json:"score"`
	Categories       []string  `json:"categories"`
}

func CreatePost(db *sql.DB, post *Post) error {
	query := `INSERT INTO posts (user_id, title, content, created_at, amount_of_comments, score) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, post.UserID, post.Title, post.Content, time.Now(), post.AmountOfComments, post.Score)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)
	return addPostCategories(db, post.ID, post.Categories)
}

func addPostCategories(db *sql.DB, postID int, categories []string) error {
	for _, category := range categories {
		var categoryID int
		err := db.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)
		if err != nil {
			return err
		}
		query := `INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)`
		_, err = db.Exec(query, postID, categoryID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPostByID(db *sql.DB, id int) (*Post, error) {
	query := `SELECT id, user_id, title, content, created_at, amount_of_comments, score FROM posts WHERE id = ?`
	row := db.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.AmountOfComments, &post.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No posts found with id %d", id)
			return nil, err
		}
		return nil, err
	}

	post.Categories, err = getPostCategories(db, post.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func getPostCategories(db *sql.DB, postID int) ([]string, error) {
	query := `SELECT c.name FROM categories c INNER JOIN post_categories pc ON c.id = pc.category_id WHERE pc.post_id = ?`
	rows, err := db.Query(query, postID)
	if err != nil {
		log.Printf("Error querying categories for post ID %d: %v", postID, err)
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			log.Printf("Error scanning category for post ID %d: %v", postID, err)
			return nil, err
		}
		categories = append(categories, category)
	}

	// debugging
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows for post id %d: %v", postID, err)
		return nil, err
	}

	log.Printf("Categories for post ID %d: %v", postID, categories)
	return categories, nil
}

func GetAllPosts(db *sql.DB) ([]Post, error) {
	query := `SELECT id, user_id, title, content, created_at, amount_of_comments, score FROM posts`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.AmountOfComments, &post.Score)
		if err != nil {
			return nil, err
		}
		post.Categories, err = getPostCategories(db, post.ID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func UpdatePostScore(db *sql.DB, postID int, score int) error {
	query := `UPDATE posts SET score = ? WHERE id = ?`
	_, err := db.Exec(query, score, postID)
	return err
}
