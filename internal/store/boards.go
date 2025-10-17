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

func DeleteBoard(id int) (bool, error) {
	db := db.GetDB()

	res, err := db.Exec("DELETE FROM boards WHERE id = $1;", id)
	if err != nil {
		fmt.Printf("Error occurred deleting board id=%d: %v\n", id, err)
		return false, err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected > 1 {
		panic(err)
	} else if rowsAffected == 0 {
		fmt.Printf("Board with id=%d not found. Unable to delete\n", id)
		return false, nil
	}

	return true, nil
}

func GetBoards() ([]models.Board, error) {
	db := db.GetDB()

	var boards []models.Board
	rows, err := db.Query("SELECT id, name, created_at FROM boards LIMIT 10;")
	if err != nil {
		fmt.Printf("Error occurred querying boards: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var board models.Board
		if err := rows.Scan(&board.Id, &board.Name, &board.CreatedAt); err != nil {
			fmt.Printf("Error occurred scanning board: %v\n", err)
			return nil, err
		}
		boards = append(boards, board)
	}

	return boards, nil
}
