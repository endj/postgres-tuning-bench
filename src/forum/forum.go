package forum

import (
	"database/sql"
	"log"

	"./post"
	"./thread"

	// fuck off
	_ "github.com/lib/pq"
)

const (
	getAllBoards    = "SELECT title FROM boards"
	getBoardThreads = "SELECT id,title,post FROM threads WHERE board_id = (SELECT id FROM boards WHERE title = $1);"
	getThread       = "SELECT threads.id, threads.title, posts.post, posts.id FROM threads, posts WHERE threads.id = $1"
)

// BoardInfo : Contains meta data about a board
type BoardInfo struct {
	Title string `json:"title"`
}

// Board : Contains meta data about the board and meta about threads
type Board struct {
	Info    BoardInfo
	Threads []thread.Info
}

// GetBoards : Fetches meta info about all boards
func GetBoards(db *sql.DB) ([]BoardInfo, error) {
	rows, err := db.Query(getAllBoards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boards := []BoardInfo{}

	for rows.Next() {
		var board BoardInfo
		err := rows.Scan(&board.Title)
		if err != nil {
			panic(err.Error())
		}
		boards = append(boards, board)
	}

	return boards, nil
}

// GetBoard : Fetches all threadsInfo and meta of a board /boards/:board"
func GetBoard(db *sql.DB, board string) (*Board, error) {
	log.Print("Got call to get board", board)

	rows, err := db.Query(getBoardThreads, board)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	threadData := []thread.Info{}

	for rows.Next() {
		var threadInfo thread.Info
		err := rows.Scan(&threadInfo.ID, &threadInfo.Title, &threadInfo.Post)
		if err != nil {
			return nil, err
		}
		threadData = append(threadData, threadInfo)
	}

	boardData := Board{
		Info: BoardInfo{
			Title: board,
		},
		Threads: threadData,
	}
	return &boardData, nil
}

// GetThread : a
func GetThread(db *sql.DB, board string, threadID int64) (*thread.Thread, error) {

	rows, err := db.Query(getThread, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []post.Post{}

	var id int
	var title string

	for rows.Next() {
		var post post.Post

		// threads.id, threads.title, posts.post, posts.id
		err := rows.Scan(&id, &title, &post.Text, &post.ID)
		if err != nil {
			return nil, err
		}
		log.Print(post.ID, post.Text, id, title)
		posts = append(posts, post)
	}
	log.Print("posts", posts)

	thread := thread.Thread{
		ID:    id,
		Title: title,
		Posts: posts,
	}
	return &thread, nil
}
