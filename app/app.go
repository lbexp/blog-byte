package app

import (
	"blog-byte/app/database"
	"blog-byte/app/middleware"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	dbConn, _ := database.Open()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("Failed to close database connection")
		}
	}()

	app := fiber.New()
	app.Use(middleware.Cors())

	validate := validator.New(validator.WithRequiredStructEnabled())

	registerRoutes(app, dbConn, validate)

	err := app.Listen(":8080")
	log.Fatal(err)
}
