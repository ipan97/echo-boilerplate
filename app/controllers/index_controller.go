package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
