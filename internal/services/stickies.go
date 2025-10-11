package services

import (
	"fmt"
	"stickies/internal/models"
	"stickies/internal/store"
	"time"
)

func CreateSticky(content string, board_id int) int {
	sticky := models.Sticky{Content: content, CreatedAt: time.Now(), BoardId: board_id}

	sticky_id, err := store.CreateSticky(&sticky)
	if err != nil {
		fmt.Printf("Error creating sticky in store: %v\n", err)
		return -1
	}

	return int(sticky_id)
}

func GetSticky(sticky_id int) *models.Sticky {
	sticky, err := store.GetSticky(sticky_id)

	if err != nil {
		fmt.Printf("Error getting sticky from store: %v\n", err)
		return nil
	}

	return sticky
}

func GetStickiesByBoard(board_id int) []models.Sticky {
	stickies, err := store.GetStickiesByBoard(board_id)

	if err != nil {
		fmt.Printf("Error getting stickies from store: %v\n", err)
		return nil
	}

	return stickies
}

func DeleteSticky(id int) bool {
	is_deleted, err := store.DeleteSticky(id)

	if err != nil {
		fmt.Printf("Error deleting sticky in store: %v\n", err)
		return false
	}

	return is_deleted
}
