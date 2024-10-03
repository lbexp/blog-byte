package user_usecase

import (
	"blog-byte/app/entity"
	user_repository "blog-byte/app/user/repository"
	bcrypt_utils "blog-byte/app/utility/bcrypt"
	error_utils "blog-byte/app/utility/error"
	jwt_utils "blog-byte/app/utility/jwt"
	"context"
	"log"
)

type UserUsecase interface {
	Login(ctx context.Context, user entity.User) (entity.User, error)
	Register(ctx context.Context, user entity.User) (entity.User, error)
}

type userUsecase struct {
	userRepo user_repository.UserRepository
}

func New(userRepo user_repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (ucase *userUsecase) Login(ctx context.Context, user entity.User) (entity.User, error) {
	userRes, err := ucase.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return entity.User{}, err
	}

	if !bcrypt_utils.CompareHashAndValue(userRes.PasswordHash, user.Password) {
		return entity.User{}, err
	}

	token, err := jwt_utils.GenerateToken(userRes)
	if err != nil {
		log.Print("User login usecase generate token error")
		return entity.User{}, error_utils.ErrorInternalServer
	}

	userRes.AccessToken = token

	return userRes, nil
}

func (ucase *userUsecase) Register(ctx context.Context, user entity.User) (entity.User, error) {
	passwordHash, err := bcrypt_utils.GenerateHash(user.Password)
	if err != nil {
		log.Print("User register usecase generate password hash error")
		return entity.User{}, error_utils.ErrorInternalServer
	}

	user.PasswordHash = passwordHash
	userRes, err := ucase.userRepo.Insert(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	token, err := jwt_utils.GenerateToken(userRes)
	if err != nil {
		log.Print("User register usecase generate token error")
		return entity.User{}, error_utils.ErrorInternalServer
	}

	userRes.AccessToken = token

	return userRes, nil
}
