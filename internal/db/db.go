package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(username, password, port string) (bool, error) {
	DB, _ := sql.Open("postgres", "postgres://"+username+":"+password+"@localhost:"+port+"/postgres?sslmode=disable")

	// Confirm a successful connection
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// Keep ref to db handle
	db = DB

	// Ensure tables exist
	if err := initializeTables(db); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func GetDB() *sql.DB {
	return db
}
