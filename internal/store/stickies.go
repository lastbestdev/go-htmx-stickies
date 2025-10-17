package store

import (
	"database/sql"
	"fmt"
	"stickies/internal/db"
	"stickies/internal/models"
)

func CreateSticky(sticky *models.Sticky) (int64, error) {
	db := db.GetDB()

	res, err := db.Exec("INSERT INTO stickies (content, board_id, color) VALUES ($1, $2, $3)", sticky.Content, sticky.BoardId, sticky.Color)

	fmt.Printf("res: %v\n", res)

	return 0, err
}

func GetSticky(sticky_id int) (*models.Sticky, error) {
	db := db.GetDB()

	var sticky models.Sticky
	err := db.QueryRow("SELECT id, content, board_id, color, created_at FROM stickies WHERE id = $1", sticky_id).Scan(
		&sticky.Id, &sticky.Content, &sticky.BoardId, &sticky.Color, &sticky.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Sticky with id=%d not found\n", sticky_id)
			return nil, err
		}
		fmt.Printf("Error occurred while attempting to retrieve Sticky w/id=%d: %v\n", sticky_id, err)
		return nil, err
	}

	return &sticky, nil
}

func GetStickiesByBoard(board_id int) ([]models.Sticky, error) {
	db := db.GetDB()

	var stickies []models.Sticky
	rows, _ := db.Query("SELECT id, content, board_id, color, created_at FROM stickies WHERE board_id = $1", board_id)

	for rows.Next() {
		var sticky models.Sticky
		err := rows.Scan(&sticky.Id, &sticky.Content, &sticky.BoardId, &sticky.Color, &sticky.CreatedAt)

		if err != nil {
			fmt.Printf("Unable to read Sticky row: %v", err)
			return nil, err
		}

		stickies = append(stickies, sticky)
	}

	return stickies, nil
}

func DeleteSticky(id int) (bool, error) {
	db := db.GetDB()

	res, err := db.Exec("DELETE FROM stickies WHERE id = $1;", id)
	if err != nil {
		fmt.Printf("Error occurred deleting sticky id=%d: %v\n", id, err)
		return false, err
	}

	delete_count, err := res.RowsAffected()

	if delete_count > 1 {
		panic(err)
	} else if delete_count == 0 {
		fmt.Printf("Sticky with id=%d not found. Unable to delete\n", id)
		return false, nil
	}

	return true, nil
}
