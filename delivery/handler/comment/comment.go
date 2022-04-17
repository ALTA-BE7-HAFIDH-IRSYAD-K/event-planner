package comment

import (
	"event-planner/delivery/middleware"
	"event-planner/delivery/response"
	"event-planner/entity"
	"event-planner/service/comment"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerComment struct {
	commentService comment.ServiceCommentInterface
}

func NewCommentHandler(commentService comment.ServiceCommentInterface) *HandlerComment {
	return &HandlerComment{
		commentService: commentService,
	}
}

func (ch *HandlerComment) CreateParticipantCommentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataComment entity.Comment

		id := middleware.ExtractToken(c)

		dataComment.UserID = uint(id)

		errBind := c.Bind(&dataComment)

		fmt.Println("data", dataComment)

		if errBind != nil {
			return errBind
		}

		err := ch.commentService.CreateComment(dataComment)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to create comment"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success to create comment"))
	}
}
