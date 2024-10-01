package app

import "github.com/gofiber/fiber/v2"

func Start() {
	app := fiber.New()

	app.Listen(":8080")
}
