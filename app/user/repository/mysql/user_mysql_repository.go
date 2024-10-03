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
	userRes := entity.User{}

	insertQuery := "INSERT INTO users(name, email, password_hash) VALUES(?, ?, ?)"

	tx, err := repo.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Print("Insert user mysql repository begin tx error: ", err)
		return userRes, err
	}

	res, err := tx.ExecContext(ctx, insertQuery, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		tx.Rollback()
		log.Print("Insert user mysql repository insert query error: ", err)
		return userRes, error_utils.ErrorInternalServer
	}

	userId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Print("Insert user mysql repository last inserted id error: ", err)
		return userRes, err
	}

	selectQuery := "SELECT * FROM users WHERE user_id = ?"

	err = tx.QueryRowContext(ctx, selectQuery, userId).Scan(
		&userRes.Id,
		&userRes.Name,
		&userRes.Email,
		&userRes.PasswordHash,
		&userRes.CreatedAt,
		&userRes.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		log.Print("Insert user query execution error: ", err)
		return userRes, error_utils.ErrorInternalServer
	}

	tx.Commit()

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
		log.Print("Select user by email query error: ", err)
		return user, error_utils.ErrorBadRequest
	}

	return user, nil
}
