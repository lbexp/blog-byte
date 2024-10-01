package post_mysql_repository

import (
	"blog-byte/app/entity"
	post_repository "blog-byte/app/post/repository"
	"context"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type postMysqlRepository struct {
	Conn *sql.DB
}

func NewPostMysqlRepository(conn *sql.DB) post_repository.PostRepository {
	return &postMysqlRepository{conn}
}

func (repo *postMysqlRepository) Insert(ctx context.Context, post entity.Post) error {
	query := "INSERT INTO posts(title, content, author_id) VALUES(?, ?, ?)"

	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Print("Insert post query preparation error")
		return err
	}

	_, err = stmt.ExecContext(ctx, post.Title, post.Content, post.AuthorId)
	if err != nil {
		log.Print("Insert post query execution error")
		return err
	}

	return nil
}

func (repo *postMysqlRepository) GetById(ctx context.Context, id int) (entity.Post, error) {
	query := "SELECT p.post_id, p.title, p.content, p.author_id, u.name as author_name, p.created_at, p.updated_at FROM posts p LEFT JOIN users u ON p.author_id = u.user_id WHERE p.author_id = ?"

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
		log.Print("Select post by ID query error")
		return post, err
	}

	return post, nil
}

func (repo *postMysqlRepository) GetAll(ctx context.Context, limit int, offset int) ([]entity.Post, error) {
	query := "SELECT p.post_id, p.title, p.content, p.author_id, u.name as author_name, p.created_at, p.updated_at FROM posts p LEFT JOIN users u ON p.author_id = u.user_id LIMIT ? OFFSET ?"

	rows, err := repo.Conn.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Print("Select posts query error")
		return nil, err
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Print("Select posts query close error")
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
			log.Print("Select posts data poopulation error")
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *postMysqlRepository) Update(ctx context.Context, post entity.Post) error {
	query := "UPDATE posts SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE post_id = ?"

	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Print("Update post query preparation error")
		return err
	}

	res, err := stmt.ExecContext(ctx, post.Title, post.Content)
	if err != nil {
		log.Print("Update post query execution error")
		return err
	}

	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		log.Print("Update post query affect no row")
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	return nil
}

func (repo *postMysqlRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM posts WHERE post_id = ?"

	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Print("Delete post query preparation error")
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		log.Print("Delete post query execution error")
		return err
	}

	affected, err := res.RowsAffected()
	if affected == 0 || err != nil {
		log.Print("Delete post query affect no row")
	}

	return nil
}
