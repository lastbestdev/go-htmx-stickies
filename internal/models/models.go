package models

import (
	"time"
)

type Board struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

type Sticky struct {
	Id        int
	BoardId   int
	Content   string
	CreatedAt time.Time
}
