package post_mysql_repository

import (
	"blog-byte/app/entity"
	post_repository "blog-byte/app/post/repository"
	error_utils "blog-byte/app/utility/error"
	"context"
	"database/sql"
	"log"
)

type postMysqlRepository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) post_repository.PostRepository {
	return &postMysqlRepository{conn}
}

func (repo *postMysqlRepository) Insert(ctx context.Context, post entity.Post) error {
	query := "INSERT INTO posts(title, content, author_id) VALUES(?, ?, ?)"

	_, err := repo.Conn.ExecContext(ctx, query, post.Title, post.Content, post.AuthorId)
	if err != nil {
		log.Print("Insert post mysql repository query execution error: ", err)
		return error_utils.ErrorInternalServer
	}

	return nil
}

func (repo *postMysqlRepository) GetById(ctx context.Context, id int) (entity.Post, error) {
	query := "SELECT p.post_id, p.title, p.content, p.author_id, u.name as author_name, p.created_at, p.updated_at FROM posts p LEFT JOIN users u ON p.author_id = u.user_id WHERE p.post_id = ?"

	post := entity.Post{}
	err := repo.Conn.QueryRowContext(ctx, query, id).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.AuthorId,
		&post.AuthorName,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		log.Print("Get by Id post mysql repository query error: ", err)
		return post, error_utils.ErrorBadRequest
	}

	return post, nil
}

func (repo *postMysqlRepository) GetAll(ctx context.Context, limit int, offset int) ([]entity.Post, error) {
	query := "SELECT p.post_id, p.title, p.content, p.author_id, u.name as author_name, p.created_at, p.updated_at FROM posts p LEFT JOIN users u ON p.author_id = u.user_id LIMIT ? OFFSET ?"

	rows, err := repo.Conn.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Print("Get all post mysql repository query error: ", err)
		return nil, error_utils.ErrorInternalServer
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Print("Get all post mysql repository query close error: ", err)
		}
	}()

	var posts []entity.Post
	for rows.Next() {
		post := entity.Post{}

		err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Print("Get all post mysql repository data population error: ", err)
			return nil, error_utils.ErrorInternalServer
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *postMysqlRepository) Update(ctx context.Context, post entity.Post) error {
	query := "UPDATE posts SET title = ?, content = ?, updated_at = now() WHERE post_id = ?"

	res, err := repo.Conn.ExecContext(ctx, query, post.Title, post.Content, post.Id)
	if err != nil {
		log.Print("Update post mysql repository query preparation error: ", err)
		return error_utils.ErrorInternalServer
	}

	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		log.Print("Update post mysql repository query affect no row: ", err)
		return error_utils.ErrorInternalServer
	}

	return nil
}

func (repo *postMysqlRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM posts WHERE post_id = ?"

	res, err := repo.Conn.ExecContext(ctx, query, id)
	if err != nil {
		log.Print("Delete post query execution error")
		return error_utils.ErrorInternalServer
	}

	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		log.Print("Delete post query affect no row")
		return error_utils.ErrorInternalServer
	}

	return nil
}
