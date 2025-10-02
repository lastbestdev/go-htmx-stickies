package main

import (
	"time"
)

type Board struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Stickys   []Sticky
}

// in-memory storage of boards (TODO: add db)
var boards = []Board{}

func createBoard(name string) Board {
	// TODO: replace with int id with guid
	id := len(boards)
	createdAt := time.Now()

	board := Board{Id: id, Name: name, Stickys: []Sticky{}, CreatedAt: createdAt}

	// TODO: update write to DB
	boards = append(boards, board)

	return board
}

func getBoard(id int) *Board {
	for i := range boards {
		if boards[i].Id == id {
			return &boards[i]
		}
	}

	return nil
}

func addSticky(board *Board, sticky Sticky) {
	board.Stickys = append(board.Stickys, sticky)
}
