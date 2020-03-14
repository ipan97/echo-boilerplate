package main

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/ipan97/echo-boilerplate/config/db"
	"github.com/ipan97/echo-boilerplate/config/di"
	"github.com/ipan97/echo-boilerplate/controllers"
	"github.com/ipan97/echo-boilerplate/models"
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

	//Set Renderer
	app.Renderer = echoview.Default()

	// Routers
	app.GET("/", controllers.Index)
	app.GET("/page", controllers.Page)

	api := app.Group("/api")
	{
		userController := di.InjectUserController()
		api.GET("/users", userController.FindAllUsers)
	}

	app.Logger.Fatal(app.Start(":9000"))

}
