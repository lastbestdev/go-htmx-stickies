package services

import (
	"fmt"
	"stickies/internal/models"
	"stickies/internal/store"
	"time"
)

func CreateBoard(name string) int {
	board := models.Board{Name: name, CreatedAt: time.Now()}

	board_id, err := store.CreateBoard(&board)
	if err != nil {
		fmt.Printf("Error creating board in store: %v\n", err)
		return -1
	}

	return int(board_id)
}

func GetBoards() []models.Board {
	boards, err := store.GetBoards()

	if err != nil {
		fmt.Printf("Error getting boards from store: %v\n", err)
		return nil
	}

	return boards
}

func GetBoard(id int) *models.Board {
	board, err := store.GetBoard(id)

	if err != nil {
		fmt.Printf("Error getting board from store: %v\n", err)
		return nil
	}

	return board
}

func DeleteBoard(id int) bool {
	is_deleted, err := store.DeleteBoard(id)

	if err != nil {
		fmt.Printf("Error deleting board in store: %v\n", err)
		return false
	}

	return is_deleted
}
