package entity

import "time"

type User struct {
	Id           int
	Name         string
	Email        string
	Password     string
	PasswordHash string
	AccessToken  string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
