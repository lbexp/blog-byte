package app

import (
	"blog-byte/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Use(middleware.Cors)

	app.Listen(":8080")
}
