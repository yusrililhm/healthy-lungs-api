package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDbInstance() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost, port=5432, username=postgres, password=")

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
