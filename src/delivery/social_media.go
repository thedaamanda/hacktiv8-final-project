package delivery

import (
	"net/http"
	"project/middleware"
	"project/src/helper"
	"project/src/model"
	"project/src/request"

	"github.com/labstack/echo"
)

type socialMediaDelivery struct {
	socialMediaUsecase model.SocialMediaUsecase
}

type SocialMediaDelivery interface {
	Mount(group *echo.Group)
}

func NewSocialMediaDelivery(SocialMediaUsecase model.SocialMediaUsecase) SocialMediaDelivery {
	return &socialMediaDelivery{
		socialMediaUsecase: SocialMediaUsecase,
	}
}

func (p *socialMediaDelivery) Mount(group *echo.Group) {
	customMiddleware := middleware.Init()
	group.Use(customMiddleware.Authentication)
	group.Use(customMiddleware.Authorization)
	group.POST("", p.StoreSocialMediaHandler)
	group.GET("", p.FetchSocialMediaHandler)
	group.PUT("/:id", p.UpdateSocialMediaHandler, customMiddleware.SocialMediaAuthorization)
	group.DELETE("/:id", p.DeleteSocialMediaHandler, customMiddleware.SocialMediaAuthorization)
}

func (p *socialMediaDelivery) StoreSocialMediaHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.SocialMediaRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	id := helper.GetUserID(c)

	socialMedia, err := p.socialMediaUsecase.CreateSocialMedia(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusCreated, "Your social media has been successfully created", socialMedia)
}

func (p *socialMediaDelivery) FetchSocialMediaHandler(c echo.Context) error {
	ctx := c.Request().Context()

	socialMedias, err := p.socialMediaUsecase.GetSocialMediaList(ctx)

	if err != nil {
		return err
	}

	return helper.SuccessResponse(c, http.StatusOK, "List of social media", map[string]interface{}{
		"social_medias": socialMedias,
	})
}

func (p *socialMediaDelivery) UpdateSocialMediaHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.SocialMediaRequest

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

	socialMedia, err := p.socialMediaUsecase.UpdateSocialMedia(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "Your social media has been successfully updated", socialMedia)
}

func (p *socialMediaDelivery) DeleteSocialMediaHandler(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := p.socialMediaUsecase.DeleteSocialMedia(ctx, id); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.MessageResponse(c, http.StatusOK, "Your social media has been successfully deleted")
}
