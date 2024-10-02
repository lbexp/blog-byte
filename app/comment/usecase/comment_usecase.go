package comment_usecase

import (
	comment_repository "blog-byte/app/comment/repository"
	"blog-byte/app/entity"
	"context"
)

type CommentUsecase interface {
	Create(ctx context.Context, comment entity.Comment) error
	GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error)
}

type commentUsecase struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentUsecase(commentRepo comment_repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}

func (ucase *commentUsecase) Create(ctx context.Context, comment entity.Comment) error {
	return ucase.commentRepo.Insert(ctx, comment)
}

func (ucase *commentUsecase) GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error) {
	return ucase.commentRepo.GetAllByPostId(ctx, postId)
}
