package controllers

import (
	"github.com/ipan97/echo-boilerplate/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	UserRepository *repository.UserRepository
}

// NewUserController create new struct
func NewUserController(userRepository *repository.UserRepository) *UserController {
	return &UserController{userRepository}
}

// FindAllUsers find all users
func (c UserController) FindAllUsers(ctx echo.Context) error {
	users, err := c.UserRepository.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, users)
}
