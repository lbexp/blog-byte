package entity

import (
	"time"
)

type Comment struct {
	Id         int
	PostId     int
	AuthorName string
	Content    string
	CreatedAt  *time.Time
}
