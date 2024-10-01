package entity

import "time"

type User struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
