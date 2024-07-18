package models

import (
	"database/sql"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateCategory(db *sql.DB, category *Category) error {
	query := `INSERT INTO categories (name) VALUES (?)`
	result, err := db.Exec(query, category.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = int(id)
	return nil
}

func GetCategoryByID(db *sql.DB, id int) (*Category, error) {
	query := `SELECT id, name FROM categories WHERE id = ?`
	row := db.QueryRow(query, id)

	var category Category
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func GetCategories(db *sql.DB) ([]Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func DeleteCategory(db *sql.DB, id int) error {
	query := `DELETE FROM categories WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

func UpdateCategory(db *sql.DB, category *Category) error {
	query := `UPDATE categories SET name = ? WHERE id = ?`
	_, err := db.Exec(query, category.Name, category.ID)
	return err
}
