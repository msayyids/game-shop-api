package controller

import (
	"net/http"

	"users/entity"
	"users/helper"

	"github.com/labstack/echo/v4"
)

func (cn *Controller) Register(c echo.Context) error {
	var reqbody entity.Users

	if err := c.Bind(&reqbody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(400, "ivnalid request"))
	}

	hashedPassword, _ := helper.HashPassword(reqbody.Password)
	reqbody.Password = hashedPassword

	user, err := cn.Controller.Adduser(reqbody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.NewErrorResponse(500, "failed to register"))
	}

	response := entity.Users{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		Phone_number: user.Phone_number,
		Age:          user.Age,
		Description:  user.Description,
		Image_url:    user.Image_url,
	}
	return c.JSON(http.StatusCreated, helper.NewResponse(201, "success register", response))
}

func (cn *Controller) Login(c echo.Context) error {

	var reqbody entity.Users

	if err := c.Bind(&reqbody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(400, "invalid request"))
	}

	loggedinUser, err := cn.Controller.FindUserbyEmail(reqbody.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, err.Error()))
	}

	validate := helper.ValidatePassword(reqbody.Password, loggedinUser.Password)
	if !validate {
		return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, "failed to login invalid password"))
	}

	strID := loggedinUser.ID.Hex()

	token, err := helper.GenerateToken(strID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.NewErrorResponse(401, "failed to create token"))
	}

	return c.JSON(http.StatusOK, helper.NewResponse(200, "success login", token))

}
