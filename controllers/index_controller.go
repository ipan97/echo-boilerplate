package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index", echo.Map{
		"title": "Index title!",
		"add": func(a int, b int) int {
			return a + b
		},
	})
}

func Page(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "page.html", echo.Map{
		"title": "Page",
	})
}
