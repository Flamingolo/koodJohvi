package handlers

import (
	"errors"
	"net/http"
	"rtf/backend/models"
)

func GetLoggedInUsedID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("RealTimeForum_session_token")
	if err != nil {
		return 0, errors.New("User not logged in")
	}

	userID, err := models.GetUserIDBySession(cookie.Value)
	if err != nil {
		return 0, errors.New("Invalid session token")
	}

	return userID, nil
}
