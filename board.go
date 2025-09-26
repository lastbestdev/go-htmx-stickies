package main

import (
	"time"
)

type Board struct {
	Id        int
	Name      string
	CreatedAt time.Time
	StickyIds []int
}

var boardCounter = 0

func createBoard(name string) Board {
	// TODO: replace with int id with guid
	id := boardCounter
	boardCounter++

	createdAt := time.Now()

	board := Board{Id: id, Name: name, StickyIds: []int{}, CreatedAt: createdAt}
	return board
}

func (board *Board) saveBoard() {
	// TODO: add save board logic here
}
