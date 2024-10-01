package post_repository

import (
	"blog-byte/app/entity"
	"context"
)

type PostRepository interface {
	Insert(ctx context.Context, post entity.Post) error
	GetById(ctx context.Context, id int) (entity.Post, error)
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Post, error)
	Update(ctx context.Context, post entity.Post) error
	Delete(ctx context.Context, id int) error
}
