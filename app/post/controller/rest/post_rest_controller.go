package post_rest_controller

import (
	"blog-byte/app/entity"
	post_controller "blog-byte/app/post/controller"
	post_usecase "blog-byte/app/post/usecase"
	error_utils "blog-byte/app/utility/error"
	jwt_utils "blog-byte/app/utility/jwt"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type postRestController struct {
	postUsecase post_usecase.PostUsecase
	validate    *validator.Validate
}

func New(postUsecase post_usecase.PostUsecase, validate *validator.Validate) post_controller.PostController {
	return &postRestController{
		postUsecase: postUsecase,
		validate:    validate,
	}
}

func (ctrl *postRestController) Create(ctx *fiber.Ctx) error {
	request := new(CreateRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		log.Print("Create post controller request body parsing error")
		return error_utils.ErrorBadRequest
	}

	err = ctrl.validate.Struct(request)
	if err != nil {
		log.Print("Create post controller not valid request error")
		return error_utils.ErrorBadRequest
	}

	jwtClaims, err := jwt_utils.GetJwtClaims(ctx)
	if err != nil {
		log.Print("Create post controller failed get jwt claims error")
		return error_utils.ErrorInternalServer
	}

	err = ctrl.postUsecase.Create(ctx.UserContext(), entity.Post{
		Title:    request.Title,
		Content:  request.Content,
		AuthorId: jwtClaims.Id,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		CreateResponse{
			Message: "Success",
		},
	)
}

func (ctrl *postRestController) GetById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Print("Delete post controller id parsing error")
		return error_utils.ErrorInternalServer
	}

	post, err := ctrl.postUsecase.GetById(ctx.UserContext(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		GetByIdResponse{
			Message: "Success",
			Data: PostDto{
				Id:         post.Id,
				Title:      post.Title,
				Content:    post.Content,
				AuthorId:   post.AuthorId,
				AuthorName: post.AuthorName,
				CreatedAt:  post.CreatedAt,
				UpdatedAt:  post.UpdatedAt,
			},
		},
	)
}

func (ctrl *postRestController) GetAll(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 20
	}

	posts, err := ctrl.postUsecase.GetAll(ctx.UserContext(), limit, page)
	if err != nil {
		return err
	}

	var data []PostDto
	for _, post := range posts {
		row := PostDto{
			Id:         post.Id,
			Title:      post.Title,
			Content:    post.Content,
			AuthorId:   post.AuthorId,
			AuthorName: post.AuthorName,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
		}

		data = append(data, row)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		GetAllResponse{
			Message: "Success",
			Data:    data,
		},
	)
}

func (ctrl *postRestController) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Print("Delete post controller id parsing error")
		return error_utils.ErrorInternalServer
	}

	request := new(UpdateRequest)
	err = ctx.BodyParser(request)
	if err != nil {
		log.Print("Update post controller request body parsing error")
		return error_utils.ErrorBadRequest
	}

	err = ctrl.validate.Struct(request)
	if err != nil {
		log.Print("Update post controller not valid request error")
		return error_utils.ErrorBadRequest
	}

	jwtClaims, err := jwt_utils.GetJwtClaims(ctx)
	if err != nil {
		log.Print("Update post controller failed get jwt claims error")
		return error_utils.ErrorInternalServer
	}

	err = ctrl.postUsecase.Update(ctx.UserContext(), entity.Post{
		Id:       id,
		Title:    request.Title,
		Content:  request.Content,
		AuthorId: jwtClaims.Id,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		UpdateResponse{
			Message: "Success",
		},
	)
}

func (ctrl *postRestController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Print("Delete post controller id parsing error")
		return error_utils.ErrorInternalServer
	}

	err = ctrl.postUsecase.Delete(ctx.UserContext(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(
		DeleteResponse{
			Message: "Success",
		},
	)
}
