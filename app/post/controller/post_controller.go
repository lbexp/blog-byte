package post_controller

import "github.com/gofiber/fiber/v2"

type PostController interface {
	Create(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
