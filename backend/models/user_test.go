package models

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateUser(t *testing.T) {
	user := &User{
		Nickname:  "testuser",
		Age:       25,
		Gender:    "other",
		FirstName: "Test",
		LastName:  "User",
		Email:     "testuser@example.com",
		Password:  "password123",
	}

	err := CreateUser(db, user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// Retrieve the user to verify it was created
	createdUser, err := GetUserByID(db, user.ID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}

	// Verify the user fields
	if createdUser.Nickname != user.Nickname {
		t.Errorf("Expected Nickname %v, got %v", user.Nickname, createdUser.Nickname)
	}
	if createdUser.Email != user.Email {
		t.Errorf("Expected Email %v, got %v", user.Email, createdUser.Email)
	}
}

func TestAuthenticateUser(t *testing.T) {
	user := &User{
		Nickname:  "authuser",
		Age:       30,
		Gender:    "female",
		FirstName: "Auth",
		LastName:  "User",
		Email:     "authuser@example.com",
		Password:  "authpassword",
	}

	err := CreateUser(db, user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// Authenticate with correct credentials
	authenticatedUser, err := AuthenticateUser(db, user.Nickname, user.Password)
	if err != nil {
		t.Fatalf("AuthenticateUser failed: %v", err)
	}
	if authenticatedUser.Nickname != user.Nickname {
		t.Errorf("Expected Nickname %v, got %v", user.Nickname, authenticatedUser.Nickname)
	}

	// Attempt to authenticate with incorrect password
	_, err = AuthenticateUser(db, user.Nickname, "wrongpassword")
	if err == nil {
		t.Fatal("Expected error for incorrect password, got none")
	}
}
