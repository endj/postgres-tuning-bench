package thread

import (
	"context"
	"database/sql"
	"log"

	"../post"
)

// Thread : struct for a thread and its posts
type Thread struct {
	ID    int
	Title string      `json:"title"`
	Posts []post.Post `json:"posts"`
}

// Info : struct for listing thread meta
type Info struct {
	ID    int
	Title string
	Post  string
}

// Input : request for creating thread
type Input struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

const (
	createPostQuery  = "INSERT INTO posts(post,thread_id) values($1,$2) returning id"
	createThreadQuery = "INSERT INTO threads(board_id, post, title) SELECT id, $1, $2 from boards where title = $3 RETURNING id"
)

// CreateThread : Creates a thread and a post linked to the thread
func CreateThread(db *sql.DB, input *Input, boardName string) (*Thread, error) {

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Print("Failed to create transaction")
		return nil, err
	}

	var threadID int
	err = tx.QueryRowContext(ctx, createThreadQuery, input.Post, input.Title, boardName).Scan(&threadID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var postID int
	err = tx.QueryRowContext(ctx, createPostQuery, input.Post, threadID).Scan(&postID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Print("Failed to commit transaction")
		return nil, err
	}

	thread := Thread{
		ID:    threadID,
		Posts: []post.Post{post.Post{ID: postID, Text: input.Post}},
		Title: input.Title,
	}
	return &thread, nil
}
