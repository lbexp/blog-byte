package user_rest_controller

import (
	"blog-byte/app/entity"
	user_controller "blog-byte/app/user/controller"
	user_usecase "blog-byte/app/user/usecase"
	error_utils "blog-byte/app/utility/error"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userRestController struct {
	userUsecase user_usecase.UserUsecase
	validate    *validator.Validate
}

func New(userUsecase user_usecase.UserUsecase, validate *validator.Validate) user_controller.UserController {
	return &userRestController{
		userUsecase: userUsecase,
		validate:    validate,
	}
}

func (ctrl *userRestController) Login(ctx *fiber.Ctx) error {
	request := new(LoginRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Print("Login controller body parsing error")
		return error_utils.ErrorBadRequest
	}

	err = ctrl.validate.Struct(request)
	if err != nil {
		log.Print("Login controller request not valid error")
		return error_utils.ErrorBadRequest
	}

	user, err := ctrl.userUsecase.Login(ctx.UserContext(), entity.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		LoginResponse{
			Message: "Success",
			Data: UserDto{
				Id:          user.Id,
				Name:        user.Name,
				Email:       user.Email,
				AccessToken: user.AccessToken,
				CreatedAt:   user.CreatedAt,
				UpdatedAt:   user.UpdatedAt,
			},
		},
	)
}

func (ctrl *userRestController) Register(ctx *fiber.Ctx) error {
	request := new(RegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Print("Register request body parsing error")
		return error_utils.ErrorBadRequest
	}

	err = ctrl.validate.Struct(request)
	if err != nil {
		log.Print("Register request not valid error")
		return error_utils.ErrorBadRequest
	}

	if request.Password != request.PasswordValidation {
		log.Print("Register request password not valid error")
		return error_utils.ErrorBadRequest
	}

	user, err := ctrl.userUsecase.Register(ctx.UserContext(), entity.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		RegisterResponse{
			Message: "Success",
			Data: UserDto{
				Id:          user.Id,
				Name:        user.Name,
				Email:       user.Email,
				AccessToken: user.AccessToken,
				CreatedAt:   user.CreatedAt,
				UpdatedAt:   user.UpdatedAt,
			},
		},
	)
}
