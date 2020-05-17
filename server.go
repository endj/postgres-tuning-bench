package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

var (
	count int    = 0
	name  string = envVariable("name")
)

func add(w http.ResponseWriter, r *http.Request) {
	count++
	log.Print("Got call to add on server ", name)
	fmt.Fprintf(w, "add %d!", count)
}

func sub(w http.ResponseWriter, r *http.Request) {
	count--
	log.Print("Got call to sub on server ", name)
	fmt.Fprintf(w, "sub %d!", count)
}

func main() {
	time.Sleep(5 * time.Second)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connectDb(psqlInfo)

	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func envVariable(key string) string {
	return os.Getenv(key)
}

func connectDb(psqlInfo string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Pinged psql!")

	sqlStatement := `INSERT INTO posts (post, replyTo) values ($1,$2)`
	_, err = db.Exec(sqlStatement, "yolo", 420)
	if err != nil {
		panic(err)
	}

	query := `SELECT post FROM posts where id=$1;`
	var post string
	row := db.QueryRow(query, 1)
	switch err := row.Scan(&post); err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		log.Println(post)
	default:
		panic(err)
	}
	defer db.Close()
}
