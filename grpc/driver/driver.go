package driver

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"strconv"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB ...
func ConnectDB() *sql.DB {
	
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	logFatal(err)
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",os.Getenv("POSTGRES_USER"),os.Getenv("POSTGRES_PASSWORD"),os.Getenv("POSTGRES_HOST"),port,os.Getenv("POSTGRES_DB"))
	pgURL, err := pq.ParseURL(url)
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	//defer db.Close()
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
