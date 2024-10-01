package comment_repository

import (
	"blog-byte/app/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) error
	GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error)
}
