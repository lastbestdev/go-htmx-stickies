package store

import (
	"stickies/internal/models"
)

func CreateSticky(sticky *models.Sticky) (int, error) {
	// TODO: implement
	return 0, nil
}

func GetSticky(sticky_id int) (*models.Sticky, error) {
	// TODO: implement
	return nil, nil
}

func GetStickiesByBoard(board_id int) ([]models.Sticky, error) {
	// TODO: implement
	return nil, nil
}

func DeleteSticky(id int) (bool, error) {
	// TODO: implement
	return false, nil
}
