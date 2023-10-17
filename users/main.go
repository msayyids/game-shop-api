package main

import (
	"users/config"
	"users/controller"
	"users/repository"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDb()
	repository := repository.NewRepository(*db)
	controller := controller.NewController(repository)

	e := echo.New()
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	// e.PUT("/user",controller.)

	e.Logger.Fatal(e.Start(":8080"))

}
