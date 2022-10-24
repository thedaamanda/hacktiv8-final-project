package delivery

import (
	"net/http"
	"project/middleware"
	"project/src/helper"
	"project/src/model"
	"project/src/request"

	"github.com/labstack/echo"
)

type commentDelivery struct {
	commentUsecase model.CommentUsecase
}

type CommentDelivery interface {
	Mount(group *echo.Group)
}

func NewCommentDelivery(CommentUsecase model.CommentUsecase) CommentDelivery {
	return &commentDelivery{
		commentUsecase: CommentUsecase,
	}
}

func (p *commentDelivery) Mount(group *echo.Group) {
	customMiddleware := middleware.Init()
	group.Use(customMiddleware.Authentication)
	group.Use(customMiddleware.Authorization)
	group.POST("", p.StoreCommentHandler)
	group.GET("", p.FetchCommentHandler)
	group.PUT("/:id", p.UpdateCommentHandler, customMiddleware.CommentAuthorization)
	group.DELETE("/:id", p.DeleteCommentHandler, customMiddleware.CommentAuthorization)
}

func (p *commentDelivery) StoreCommentHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.CreateCommentRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	id := helper.GetUserID(c)

	comment, err := p.commentUsecase.CreateComment(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusCreated, "Your comment has been successfully created", comment)
}

func (p *commentDelivery) FetchCommentHandler(c echo.Context) error {
	ctx := c.Request().Context()

	comments, err := p.commentUsecase.GetCommentList(ctx)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "List of comments", comments)
}

func (p *commentDelivery) UpdateCommentHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.UpdateCommentRequest

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	comment, err := p.commentUsecase.UpdateComment(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "Your comment has been successfully updated", comment)
}

func (p *commentDelivery) DeleteCommentHandler(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := p.commentUsecase.DeleteComment(ctx, id); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.MessageResponse(c, http.StatusOK, "Your comment has been successfully deleted")
}
