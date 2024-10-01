package entity

import "time"

type Post struct {
	Id         int
	Title      string
	Content    string
	AuthorId   int
	AuthorName string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
