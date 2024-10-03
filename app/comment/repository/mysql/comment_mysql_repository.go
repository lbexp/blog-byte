package comment_mysql_repository

import (
	comment_repository "blog-byte/app/comment/repository"
	"blog-byte/app/entity"
	error_utils "blog-byte/app/utility/error"
	"context"
	"database/sql"
	"log"
)

type commentMysqlRepository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) comment_repository.CommentRepository {
	return &commentMysqlRepository{conn}
}

func (repo *commentMysqlRepository) Insert(ctx context.Context, comment entity.Comment) error {
	query := "INSERT INTO comments(post_id, author_name, content) VALUES(?, ?, ?)"

	_, err := repo.Conn.ExecContext(ctx, query, comment.PostId, comment.AuthorName, comment.Content)
	if err != nil {
		log.Print("Insert comment mysql repository query execution error: ", err)
		return error_utils.ErrorInternalServer
	}

	return nil
}

func (repo *commentMysqlRepository) GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error) {
	query := "SELECT comment_id, post_id, author_name, content, created_at FROM comments WHERE post_id = ?"

	rows, err := repo.Conn.QueryContext(ctx, query, postId)
	if err != nil {
		log.Print("Get all by post id comment repository query error: ", err)
		return nil, error_utils.ErrorInternalServer
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Print("Get all by post id comment repository query close error: ", err)
		}
	}()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}

		err = rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.AuthorName,
			&comment.Content,
			&comment.CreatedAt,
		)
		if err != nil {
			log.Print("Get all by post id comment repository data population error: ", err)
			return nil, error_utils.ErrorInternalServer
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
