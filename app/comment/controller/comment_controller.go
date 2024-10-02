package comment_controller

import "github.com/gofiber/fiber/v2"

type CommentController interface {
	Create(ctx *fiber.Ctx) error
	GetAllByPostId(ctx *fiber.Ctx) error
}
