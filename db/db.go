package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDataBase() *sql.DB {
	connection := "user=postgres dbname=go-api password=senha host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
