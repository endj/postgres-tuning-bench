package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	driver string = "postgres"
)

var (
	host     string = os.Getenv("PG_HOST")
	port     int    = getEnvAsInt("PG_PORT")
	user     string = os.Getenv("PG_USER")
	password string = os.Getenv("PG_PASSWORD")
	dbname   string = os.Getenv("PG_DB_NAME")
)

func GetDataBaseConnection() *sql.DB {
	properties := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	retries := 3
	for retries > 0 {
		db, err := sql.Open(driver, properties)
		if err == nil {
			log.Print("Connected to DB")
			return db
		}
		time.Sleep(5000 * time.Millisecond)
		retries = retries - 1
	}
	panic("Failed to connect to DB in 15 seconds")
}

func getEnvAsInt(key string) int {
	number, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return number
}
