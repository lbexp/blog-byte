package comment_mysql_repository

import (
	comment_repository "blog-byte/app/comment/repository"
	"blog-byte/app/entity"
	"context"
	"database/sql"
	"log"
)

type commentMysqlRepository struct {
	Conn *sql.DB
}

func NewCommentMysqlRepository(conn *sql.DB) comment_repository.CommentRepository {
	return &commentMysqlRepository{conn}
}

func (repo *commentMysqlRepository) Insert(ctx context.Context, comment entity.Comment) error {
	query := "INSERT INTO comments(post_id, author_name, content) VALUES(?, ?, ?)"

	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Print("Insert comment query preparation error")
		return err
	}

	_, err = stmt.ExecContext(ctx, comment.PostId, comment.AuthorName, comment.Content)
	if err != nil {
		log.Print("Insert comment query execution error")
		return err
	}

	return nil
}

func (repo *commentMysqlRepository) GetAllByPostId(ctx context.Context, postId int) ([]entity.Comment, error) {
	query := "SELECT comment_id, post_id, author_name, content, created_at FROM comments WHERE post_id = ?"

	rows, err := repo.Conn.QueryContext(ctx, query, postId)
	if err != nil {
		log.Print("Select comments by post_id query error")
		return nil, err
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Print("Select comments by post_id query close error")
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
			log.Print("Select comments by post_id data population error")
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
