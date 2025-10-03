package services

import (
	"time"
)

type Board struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

// in-memory storage of boards (TODO: add db)
var boards = []Board{}

func CreateBoard(name string) Board {
	// TODO: replace with int id with guid
	id := len(boards)
	createdAt := time.Now()

	board := Board{Id: id, Name: name, CreatedAt: createdAt}

	// TODO: update write to DB
	boards = append(boards, board)

	return board
}

func GetBoard(id int) *Board {
	for i := range boards {
		if boards[i].Id == id {
			return &boards[i]
		}
	}

	return nil
}
