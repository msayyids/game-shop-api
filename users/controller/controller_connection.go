package controller

import "users/repository"

type Controller struct {
	Controller repository.Repository
}

func NewController(repo repository.Repository) Controller {
	return Controller{Controller: repo}
}
