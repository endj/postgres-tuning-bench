package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"os"

	"./db"
	"./forum"
	"./forum/post"
	"./forum/thread"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var (
	connection = *db.GetDataBaseConnection()
)

func getBoards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boards, err := forum.GetBoards(&connection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json, err := json.Marshal(boards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", string(json))
}

func getBoard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boardParam := ps.ByName("board")
	if boardParam == "" {
		http.Error(w, "Board not found", 404)
		return
	}
	board, err := forum.GetBoard(&connection, boardParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", string(json))
}

func getThread(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boardParam := ps.ByName("board")
	threadParam := ps.ByName("thread")
	if boardParam == "" || threadParam == "" {
		http.Error(w, "Thread not found", 404)
		return
	}

	threadID, err := strconv.ParseInt(threadParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid thread", 400)
		return
	}
	thread, err := forum.GetThread(&connection, boardParam, threadID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(thread)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", string(json))
}

func createThread(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boardParam := ps.ByName("board")
	if boardParam == "" {
		http.Error(w, "Board not found", 404)
		return
	}
	log.Print("Create Thread BoardParam", boardParam)

	var request thread.Input
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	thread, err := thread.CreateThread(&connection, &request, boardParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json, err := json.Marshal(thread)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", string(json))
}

func createReply(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boardParam := ps.ByName("board")
	threadParam := ps.ByName("thread")
	if boardParam == "" || threadParam == "" {
		http.Error(w, "Thread not found", 404)
		return
	}

	threadID, err := strconv.ParseInt(threadParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid thread", 400)
		return
	}

	var request post.Input
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	post, err := post.CreatePost(&connection, &request, threadID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", string(json))
}

func handleRootCall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "https://www.localhost/boards", 302)
}

func main() {
	router := httprouter.New()
	router.GET("/", handleRootCall)
	router.GET("/boards", getBoards)
	router.GET("/boards/:board", getBoard)
	router.GET("/boards/:board/threads/:thread", getThread)
	router.POST("/boards/:board/threads", createThread)
	router.POST("/boards/:board/threads/:thread/reply", createReply)

	log.Print("Starting server ", os.Getenv("name"))

	defer connection.Close()
	log.Fatal(http.ListenAndServe(":8080", router))
}
