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

var boardCounter = 0

func createBoard(name string) Board {
	// TODO: replace with int id with guid
	id := boardCounter
	boardCounter++

	createdAt := time.Now()

	board := Board{Id: id, Name: name, Stickys: []Sticky{}, CreatedAt: createdAt}
	return board
}

func addSticky(board *Board, sticky Sticky) {
	board.Stickys = append(board.Stickys, sticky)
}

func (board *Board) saveBoard() {
	// TODO: add save board logic here
}
