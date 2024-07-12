BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"nickname" TEXT UNIQUE,
"age" INTEGER,
"gender" TEXT,
"first_name" TEXT,
"last_name" TEXT,
"email" TEXT UNIQUE,
"password" TEXT,
"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
"post_count" INTEGER NOT NULL DEFAULT 0,
"comment_count" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "active_sessions" (
    "user_id" INTEGER NOT NULL,
    "session_id" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "expires_at" TIMESTAMP NOT NULL,
    "last_activity" TIMESTAMP NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);

CREATE TABLE IF NOT EXISTS "posts" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER,
    "title" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "amount_of_comments" INTEGER NOT NULL DEFAULT 0,
    "score" INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "categories" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "post_categories" (
    "post_id" INTEGER,
    "category_id" INTEGER,
    PRIMARY KEY("post_id", "category_id"),
    FOREIGN KEY("post_id") REFERENCES "posts"("id"),
    FOREIGN KEY("category_id") REFERENCES "categories"("id")
);
CREATE TABLE IF NOT EXISTS "comments" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "post_id" INTEGER,
    "user_id" INTEGER,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "score" INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY("post_id") REFERENCES "posts"("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "scores" (
    "user_id" INTEGER NOT NULL,
    "post_id" INTEGER,
    "comment_id" INTEGER,
    "status" STRING NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users"("id"),
    FOREIGN KEY("post_id") REFERENCES "posts"("id"),
    FOREIGN KEY("comment_id") REFERENCES "comments"("id")
);
CREATE TABLE IF NOT EXISTS "messages" (
    "sender_id" INTEGER NOT NULL,
    "receiver_id" INTEGER NOT NULL,
    "content" TEXT NOT NULL,
    "is_read" BOOLEAN,
    FOREIGN KEY("sender_id") REFERENCES "users"("id"),
    FOREIGN KEY("receiver_id") REFERENCES "users"("id")
);
COMMIT;