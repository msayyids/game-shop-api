package middleware

import (
	"net/http"
	"users/config"
	"users/helper"
	"users/repository"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		repo := repository.NewRepository(*config.ConnectDb())
		token := c.Request().Header.Get("token")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, "invalid access token"))
		}

		claims, err := helper.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, "invalid access token"))
		}

		id := claims["id"].(string)
		loggedinUser, err := repo.FindUserById(id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, "invalid access token"))
		}

		c.Set("loggedInUser", loggedinUser)

		return next(c)

	}
}
