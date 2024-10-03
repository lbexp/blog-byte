package comment_rest_controller

import (
	comment_controller "blog-byte/app/comment/controller"
	comment_usecase "blog-byte/app/comment/usecase"
	"blog-byte/app/entity"
	error_utils "blog-byte/app/utility/error"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type commentRestController struct {
	commentUsecase comment_usecase.CommentUsecase
	validate       *validator.Validate
}

func New(commentUsecase comment_usecase.CommentUsecase, validate *validator.Validate) comment_controller.CommentController {
	return &commentRestController{
		commentUsecase: commentUsecase,
		validate:       validate,
	}
}

func (ctrl *commentRestController) Create(ctx *fiber.Ctx) error {
	postId, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		log.Print("Create comment controller post_id parsing error: ", err)
		return error_utils.ErrorInternalServer
	}

	request := new(CreateRequest)
	err = ctx.BodyParser(request)
	if err != nil {
		log.Print("Create comment controller body parsing error: ", err)
		return error_utils.ErrorBadRequest
	}

	err = ctrl.validate.Struct(request)
	if err != nil {
		log.Print("Create comment controller not valid request error: ", err)
		return error_utils.ErrorBadRequest
	}

	err = ctrl.commentUsecase.Create(ctx.UserContext(), entity.Comment{
		PostId:     postId,
		Content:    request.Content,
		AuthorName: request.AuthorName,
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

func (ctrl *commentRestController) GetAllByPostId(ctx *fiber.Ctx) error {
	postId, err := strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		log.Print("Delete post controller id parsing error: ", err)
		return error_utils.ErrorInternalServer
	}

	comments, err := ctrl.commentUsecase.GetAllByPostId(ctx.UserContext(), postId)
	if err != nil {
		return err
	}

	var data []CommentDto
	for _, comment := range comments {
		row := CommentDto{
			Id:         comment.Id,
			PostId:     comment.PostId,
			AuthorName: comment.AuthorName,
			Content:    comment.Content,
			CreatedAt:  comment.CreatedAt,
		}

		data = append(data, row)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		GetAllByPostIdResponse{
			Message: "Success",
			Data:    data,
		},
	)
}
