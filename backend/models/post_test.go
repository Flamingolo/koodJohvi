package models

import (
    "testing"
)

func TestCreatePost(t *testing.T) {
    user := &User{
        Nickname:  "postuser",
        Age:       25,
        Gender:    "other",
        FirstName: "Post",
        LastName:  "User",
        Email:     "postuser@example.com",
        Password:  "password123",
    }
    err := CreateUser(db, user)
    if err != nil {
        t.Fatalf("CreateUser failed: %v", err)
    }

    post := &Post{
        UserID:   user.ID,
        Category: "General",
        Title:    "Test Post",
        Content:  "This is a test post.",
        Likes:    0,
        Dislikes: 0,
    }

    err = CreatePost(db, post)
    if err != nil {
        t.Fatalf("CreatePost failed: %v", err)
    }

    // Retrieve the post to verify it was created
    createdPost, err := GetPostByID(db, post.ID)
    if err != nil {
        t.Fatalf("GetPostByID failed: %v", err)
    }

    // Verify the post fields
    if createdPost.Title != post.Title {
        t.Errorf("Expected Title %v, got %v", post.Title, createdPost.Title)
    }
    if createdPost.Content != post.Content {
        t.Errorf("Expected Content %v, got %v", post.Content, createdPost.Content)
    }
}

func TestGetPosts(t *testing.T) {
    user := &User{
        Nickname:  "postuser2",
        Age:       26,
        Gender:    "male",
        FirstName: "Post2",
        LastName:  "User2",
        Email:     "postuser2@example.com",
        Password:  "password456",
    }
    err := CreateUser(db, user)
    if err != nil {
        t.Fatalf("CreateUser failed: %v", err)
    }

    post := &Post{
        UserID:   user.ID,
        Category: "General",
        Title:    "Test Post 2",
        Content:  "This is another test post.",
        Likes:    0,
        Dislikes: 0,
    }

    err = CreatePost(db, post)
    if err != nil {
        t.Fatalf("CreatePost failed: %v", err)
    }

    posts, err := GetPosts(db)
    if err != nil {
        t.Fatalf("GetPosts failed: %v", err)
    }

    if len(posts) == 0 {
        t.Fatalf("Expected at least one post, got %d", len(posts))
    }
}

func TestUpdatePost(t *testing.T) {
    post := &Post{
        UserID:   1,
        Category: "Update",
        Title:    "Initial Title",
        Content:  "Initial Content",
        Likes:    0,
        Dislikes: 0,
    }

    err := CreatePost(db, post)
    if err != nil {
        t.Fatalf("CreatePost failed: %v", err)
    }

    post.Title = "Updated Title"
    post.Content = "Updated Content"

    err = UpdatePost(db, post)
    if err != nil {
        t.Fatalf("UpdatePost failed: %v", err)
    }

    updatedPost, err := GetPostByID(db, post.ID)
    if err != nil {
        t.Fatalf("GetPostByID failed: %v", err)
    }

    if updatedPost.Title != post.Title {
        t.Errorf("Expected Title %v, got %v", post.Title, updatedPost.Title)
    }
    if updatedPost.Content != post.Content {
        t.Errorf("Expected Content %v, got %v", post.Content, updatedPost.Content)
    }
}

func TestDeletePost(t *testing.T) {
    post := &Post{
        UserID:   1,
        Category: "Delete",
        Title:    "To Be Deleted",
        Content:  "This post will be deleted.",
        Likes:    0,
        Dislikes: 0,
    }

    err := CreatePost(db, post)
    if err != nil {
        t.Fatalf("CreatePost failed: %v", err)
    }

    err = DeletePost(db, post.ID)
    if err != nil {
        t.Fatalf("DeletePost failed: %v", err)
    }

    _, err = GetPostByID(db, post.ID)
    if err == nil {
        t.Fatal("Expected error for non-existent post, got none")
    }
}
