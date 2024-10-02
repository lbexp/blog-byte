package user_repository

import (
	"blog-byte/app/entity"
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}
