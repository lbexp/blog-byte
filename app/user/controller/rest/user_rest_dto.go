package user_rest_controller

import "time"

type UserDto struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	AccessToken string     `json:"access_token"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type LoginResponse struct {
	Message string  `json:"message"`
	Data    UserDto `json:"data"`
}

type RegisterRequest struct {
	Name               string `json:"name" validate:"required,min=6,max=50"`
	Email              string `json:"email" validate:"required,email"`
	Password           string `json:"password" validate:"required,min=6,max=50"`
	PasswordValidation string `json:"password_validation" validate:"required,min=6,max=50"`
}

type RegisterResponse struct {
	Message string  `json:"message"`
	Data    UserDto `json:"data"`
}
