package app

import (
	user_rest_controller "blog-byte/app/user/controller/rest"
	user_mysql_repository "blog-byte/app/user/repository/mysql"
	user_usecase "blog-byte/app/user/usecase"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App, dbConn *sql.DB, validate *validator.Validate) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	userRepo := user_mysql_repository.NewUserMysqlRepository(dbConn)
	userUcase := user_usecase.NewUserUsecase(userRepo)
	userCtrl := user_rest_controller.NewUserRestController(userUcase, validate)

	auth := v1.Group("/auth")
	auth.Post("/login", userCtrl.Login)
	auth.Post("/register", userCtrl.Register)
}
