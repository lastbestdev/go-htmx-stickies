package store

import (
	"database/sql"
	"fmt"
	"stickies/internal/db"
	"stickies/internal/models"
)

func CreateBoard(board *models.Board) (int64, error) {
	db := db.GetDB()

	res, err := db.Exec("INSERT INTO boards (name) VALUES ($1);", board.Name)

	fmt.Printf("res: %v\n", res)

	return 0, err
}

func GetBoard(id int) (*models.Board, error) {
	db := db.GetDB()

	var board models.Board
	err := db.QueryRow("SELECT id, name, created_at FROM boards WHERE id = $1", id).Scan(
		&board.Id, &board.Name, &board.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No board found for id=%d\n", id)
			return nil, err
		}
		fmt.Printf("Error occurred querying boards for id=%d: %v\n", id, err)
		return nil, err
	}

	return &board, nil
}
