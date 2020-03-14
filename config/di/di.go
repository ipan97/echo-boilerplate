package di

import (
	"github.com/ipan97/echo-boilerplate/config/db"
	"github.com/ipan97/echo-boilerplate/controllers"
	"github.com/ipan97/echo-boilerplate/repository"
)

func InjectUserController() *controllers.UserController {
	userRepository := repository.NewUserRepository(db.DB())
	userController := controllers.NewUserController(userRepository)
	return userController
}
