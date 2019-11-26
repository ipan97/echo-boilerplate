package main

import (
	"github.com/ipan97/echo-boilerplate/app/controllers"
	"github.com/ipan97/echo-boilerplate/app/models"
	"github.com/ipan97/echo-boilerplate/conf/db"
	"github.com/ipan97/echo-boilerplate/conf/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	db.DB().AutoMigrate(models.User{})
}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routers
	app.GET("/", controllers.Index)

	api := app.Group("/api")
	{
		userController := di.InjectUserController()
		api.GET("/users", userController.FindAllUsers)
	}

	app.Logger.Fatal(app.Start(":9000"))

}
