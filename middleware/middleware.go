package middleware

import (
	"net/http"
	"project/config/postgres"
	"project/helper/token"
	"project/src/helper"
	"project/src/model"
	"strings"

	"github.com/labstack/echo"
)

type Middleware struct{}

func Init() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if !strings.Contains(tokenString, "Bearer") {
			return helper.ErrorResponse(c, http.StatusUnauthorized, "Token not provided")
		}

		tokenStrings := strings.Replace(tokenString, "Bearer ", "", 1)

		extract, err := token.ExtractToken(tokenStrings, token.TypeShortSecretKey)

		if err != nil {
			return helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set("userData", extract)
		return next(c)
	}
}

func (m *Middleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User

		userID := helper.GetUserID(c)

		err := postgres.InitGorm().Where("id = ?", userID).First(&user).Error

		if err != nil {
			return helper.ErrorResponse(c, http.StatusForbidden, "User does not exist")
		}

		return next(c)
	}
}

func (m *Middleware) PhotoAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var photo model.Photo

		photoID := helper.GetID(c)

		if err := postgres.InitGorm().Where("id = ?", photoID).First(&photo).Error; err != nil {
			return helper.ErrorResponse(c, http.StatusNotFound, "Photo does not exist")
		}

		if photo.UserID != helper.GetUserID(c) {
			return helper.ErrorResponse(c, http.StatusForbidden, "You are not authorized to access this photo")
		}

		return next(c)
	}
}

func (m *Middleware) CommentAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var comment model.Comment

		commentID := helper.GetID(c)

		if err := postgres.InitGorm().Where("id = ?", commentID).First(&comment).Error; err != nil {
			return helper.ErrorResponse(c, http.StatusNotFound, "Comment does not exist")
		}

		if comment.UserID != helper.GetUserID(c) {
			return helper.ErrorResponse(c, http.StatusForbidden, "You are not authorized to access this comment")
		}

		return next(c)
	}
}

func (m *Middleware) SocialMediaAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var socialMedia model.SocialMedia

		socialMediaID := helper.GetID(c)

		if err := postgres.InitGorm().Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
			return helper.ErrorResponse(c, http.StatusNotFound, "Social Media does not exist")
		}

		if socialMedia.UserID != helper.GetUserID(c) {
			return helper.ErrorResponse(c, http.StatusForbidden, "You are not authorized to access this social media")
		}

		return next(c)
	}
}
