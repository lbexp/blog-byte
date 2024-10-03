package comment_rest_controller

import "time"

type CommentDto struct {
	Id         int        `json:"id"`
	PostId     int        `json:"post_id"`
	AuthorName string     `json:"author_name"`
	Content    string     `json:"content"`
	CreatedAt  *time.Time `json:"created_at"`
}

type CreateRequest struct {
	AuthorName string `json:"author_name" validate:"required,min=3,max=50"`
	Content    string `json:"content" validate:"required"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type GetAllByPostIdResponse struct {
	Message string       `json:"message"`
	Data    []CommentDto `json:"data"`
}
