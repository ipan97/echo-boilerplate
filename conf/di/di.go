package di

import (
	"github.com/ipan97/echo-boilerplate/app/controllers"
	"github.com/ipan97/echo-boilerplate/app/repository"
	"github.com/ipan97/echo-boilerplate/conf/db"
)

func InjectUserController() *controllers.UserController {
	userRepository := repository.NewUserRepository(db.DB())
	userController := controllers.NewUserController(userRepository)
	return userController
}
