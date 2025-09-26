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
}

var stickyCounter = 0

func createSticky(content string, x int, y int) Sticky {
	// TODO: replace with int id with guid
	id := stickyCounter
	stickyCounter++

	createdAt := time.Now()

	// create sticky note
	sticky := Sticky{Id: id, X: x, Y: y, Content: content, CreatedAt: createdAt}

	return sticky
}
