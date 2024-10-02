package post_usecase

import (
	"blog-byte/app/entity"
	post_repository "blog-byte/app/post/repository"
	"context"
)

type PostUsecase interface {
	Create(ctx context.Context, post entity.Post) error
	GetById(ctx context.Context, id int) (entity.Post, error)
	GetAll(ctx context.Context, limit int, offset int) ([]entity.Post, error)
	Update(ctx context.Context, post entity.Post) error
	Delete(ctx context.Context, id int) error
}

type postUsecase struct {
	postRepo post_repository.PostRepository
}

func New(postRepo post_repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepo: postRepo,
	}
}

func (ucase *postUsecase) Create(ctx context.Context, post entity.Post) error {
	return ucase.postRepo.Insert(ctx, post)
}

func (ucase *postUsecase) GetById(ctx context.Context, id int) (entity.Post, error) {
	return ucase.postRepo.GetById(ctx, id)
}

func (ucase *postUsecase) GetAll(ctx context.Context, limit int, offest int) ([]entity.Post, error) {
	return ucase.postRepo.GetAll(ctx, limit, offest)
}

func (ucase *postUsecase) Update(ctx context.Context, post entity.Post) error {
	return ucase.postRepo.Update(ctx, post)
}

func (ucase *postUsecase) Delete(ctx context.Context, id int) error {
	return ucase.postRepo.Delete(ctx, id)
}
