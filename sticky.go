package main

import (
	"time"
)

type Sticky struct {
	Id        int
	X         int
	Y         int
	Content   string
	CreatedAt time.Time
	Board     Board
}

var stickyCounter = 0

func createSticky(content string, x int, y int, board Board) *Sticky {
	// TODO: replace with int id with guid
	id := stickyCounter
	stickyCounter++

	createdAt := time.Now()

	return &Sticky{Id: id, X: x, Y: y, Content: content, Board: board, CreatedAt: createdAt}
}
