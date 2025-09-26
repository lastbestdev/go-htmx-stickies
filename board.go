package main

import (
	"time"
)

type Board struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Stickies  []Sticky
}

var boardCounter = 0

func createBoard(name string) *Board {
	// TODO: replace with int id with guid
	id := boardCounter
	boardCounter++

	createdAt := time.Now()

	return &Board{Id: id, Name: name, Stickies: []Sticky{}, CreatedAt: createdAt}
}

func (board *Board) saveBoard() {
	// TODO: add save board logic here
}
