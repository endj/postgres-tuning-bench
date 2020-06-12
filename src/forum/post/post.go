package post

import "database/sql"

// Post : A struct for post field
type Post struct {
	ID   int
	Text string
}

// Input : input for creating a reply
type Input struct {
	Text string
}

const (
	createPostQuery = "INSERT INTO posts(post,thread_id) values($1,$2) returning id"
)

// CreatePost : creates a post
func CreatePost(db *sql.DB, input *Input, threadID int64) (*Post, error) {

	var postID int
	err := db.QueryRow(createPostQuery, input.Text, threadID).Scan(&postID)
	if err != nil {
		return nil, err
	}
	post := Post{
		ID:   postID,
		Text: input.Text,
	}
	return &post, nil
}
