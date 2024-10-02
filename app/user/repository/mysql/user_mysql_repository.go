package user_mysql_repository

import (
	"blog-byte/app/entity"
	user_repository "blog-byte/app/user/repository"
	error_utils "blog-byte/app/utility/error"
	"context"
	"database/sql"
	"log"
)

type userMysqlRepository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) user_repository.UserRepository {
	return &userMysqlRepository{conn}
}

func (repo *userMysqlRepository) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	query := "INSERT INTO users(name, email, password_hash) VALUES(?, ?, ?) RETURNING *"

	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Print("Insert user query preparation error")
		return entity.User{}, error_utils.ErrorInternalServer
	}

	userRes := entity.User{}
	err = stmt.QueryRowContext(ctx, user.Name, user.Email, user.PasswordHash).Scan(
		&userRes.Id,
		&userRes.Name,
		&userRes.Email,
		&userRes.PasswordHash,
		&userRes.CreatedAt,
		&userRes.UpdatedAt,
	)
	if err != nil {
		log.Print("Insert user query execution error")
		return entity.User{}, error_utils.ErrorInternalServer
	}

	return userRes, nil
}

func (repo *userMysqlRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	query := "SELECT user_id, name, email, password_hash, created_at, updated_at FROM users WHERE email = ?"

	user := entity.User{}
	err := repo.Conn.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Print("Select user by email query error")
		return user, error_utils.ErrorInternalServer
	}

	return user, nil
}
