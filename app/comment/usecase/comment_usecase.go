package comment_usecase

import (
	comment_repository "blog-byte/app/comment/repository"
	"blog-byte/app/entity"
	post_repository "blog-byte/app/post/repository"
	"context"
)

type CommentUsecase interface {
	Create(ctx context.Context, comment entity.Comment) error
	GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error)
}

type commentUsecase struct {
	commentRepo comment_repository.CommentRepository
	postRepo    post_repository.PostRepository
}

func NewCommentUsecase(commentRepo comment_repository.CommentRepository, postRepo post_repository.PostRepository) CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepo,
	}
}

func (ucase *commentUsecase) Create(ctx context.Context, comment entity.Comment) error {
	_, err := ucase.postRepo.GetById(ctx, comment.PostId)
	if err != nil {
		return err
	}

	return ucase.commentRepo.Insert(ctx, comment)
}

func (ucase *commentUsecase) GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error) {
	_, err := ucase.postRepo.GetById(ctx, postId)
	if err != nil {
		return nil, err
	}

	return ucase.commentRepo.GetAllByPostId(ctx, postId)
}
