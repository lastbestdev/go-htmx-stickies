package services

import (
	"time"
)

type Sticky struct {
	Id        int
	X         int
	Y         int
	Content   string
	CreatedAt time.Time
	Board     *Board
}

// in-memory storage of stickies (TODO: add db)
var stickies = []Sticky{}

func CreateSticky(content string, board *Board) Sticky {
	// TODO: replace with int id with guid
	id := len(stickies)

	createdAt := time.Now()

	// create sticky note
	sticky := Sticky{Id: id, Content: content, CreatedAt: createdAt, Board: board}

	// add to in-memory store (TODO: replace with db)
	stickies = append(stickies, sticky)

	return sticky
}

func GetStickies(Board Board) []Sticky {
	var results []Sticky

	for i := range stickies {
		if stickies[i].Board != nil && stickies[i].Board.Id == Board.Id {
			results = append(results, stickies[i])
		}
	}

	return results
}
