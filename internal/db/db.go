package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(username, password, port string) (bool, error) {
	db, _ := sql.Open("postgres", "postgres://"+username+":"+password+"@localhost:"+port+"/postgres?sslmode=disable")

	// Confirm a successful connection.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Ensure tables exist
	if err := initializeTables(db); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
