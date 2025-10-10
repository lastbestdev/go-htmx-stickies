package db

import (
	"context"
	"database/sql"
	"time"
)

func initializeTables(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS boards (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS stickies (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			board_id INTEGER REFERENCES boards(id) ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	return nil
}
