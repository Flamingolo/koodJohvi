package models

import "database/sql"

func CreateSessions(token string, userID int) error {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO sessions (token, user_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(token, userID)
	return err
}

func GetUserIDBySession(token string) (int, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var userID int
	err = db.QueryRow("SELECT user_id FROM sessions WHERE token = ?").Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
