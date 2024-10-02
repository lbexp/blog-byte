package app

import (
	comment_rest_controller "blog-byte/app/comment/controller/rest"
	comment_mysql_repository "blog-byte/app/comment/repository/mysql"
	comment_usecase "blog-byte/app/comment/usecase"
	post_rest_controller "blog-byte/app/post/controller/rest"
	post_mysql_repository "blog-byte/app/post/repository/mysql"
	post_usecase "blog-byte/app/post/usecase"
	user_rest_controller "blog-byte/app/user/controller/rest"
	user_mysql_repository "blog-byte/app/user/repository/mysql"
	user_usecase "blog-byte/app/user/usecase"
	"database/sql"
	"os"

	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App, dbConn *sql.DB, validate *validator.Validate) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	posts := v1.Group("/posts")
	postId := posts.Group("/:post_id")
	comments := postId.Group("/comments")

	userRepo := user_mysql_repository.NewUserMysqlRepository(dbConn)
	userUcase := user_usecase.NewUserUsecase(userRepo)
	userCtrl := user_rest_controller.NewUserRestController(userUcase, validate)

	postRepo := post_mysql_repository.NewPostMysqlRepository(dbConn)
	postUcase := post_usecase.NewPostUsecase(postRepo)
	postCtrl := post_rest_controller.NewPostRestController(postUcase, validate)

	commentRepo := comment_mysql_repository.NewCommentMysqlRepository(dbConn)
	commentUcase := comment_usecase.NewCommentUsecase(commentRepo)
	commentCtrl := comment_rest_controller.NewCommentRestController(commentUcase, validate)

	// Public routes - begin
	v1.Post("/login", userCtrl.Login)
	v1.Post("/register", userCtrl.Register)

	posts.Get("/", postCtrl.GetAll)
	postId.Get("/", postCtrl.GetById)

	comments.Get("/", commentCtrl.GetAllByPostId)
	// Public routes - end

	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_KEY"))},
	}))

	// Protected routes - begin
	posts.Post("/", postCtrl.Create)
	postId.Put("/", postCtrl.Update)
	postId.Delete("/", postCtrl.Delete)

	comments.Post("/", commentCtrl.Create)
	// Protected routes - end
}
