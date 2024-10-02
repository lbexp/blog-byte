package post_rest_controller

import "time"

type PostDto struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	AuthorId   int        `json:"author_id"`
	AuthorName string     `json:"author_name"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type GetByIdResponse struct {
	Message string  `json:"message"`
	Data    PostDto `json:"data"`
}

type GetAllResponse struct {
	Message string    `json:"message"`
	Data    []PostDto `json:"data"`
}

type CreateRequest struct {
	Title   string `json:"title" validate:"required,max=255"`
	Content string `json:"content" validate:"required"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type UpdateRequest struct {
	Title   string `json:"title" validate:"required,max=255"`
	Content string `json:"content" validate:"required"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
