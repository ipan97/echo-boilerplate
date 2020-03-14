package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Meta struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type Data interface {
}

type Errors interface {
}

type ApiResponse struct {
	Meta   `json:"meta"`
	Data   `json:"data"`
	Errors `json:"errors"`
}

func Ok(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, ApiResponse{
		Meta: Meta{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
		},
		Data: data,
	})
}

func InternalServerError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, ApiResponse{
		Meta: Meta{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
		},
		Errors: err.Error(),
	})
}

func NotFound(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusNotFound, ApiResponse{
		Meta: Meta{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
		},
		Errors: err.Error(),
	})
}

func Unauthorized(ctx echo.Context, errorMessage string) error {
	return ctx.JSON(http.StatusUnauthorized, ApiResponse{
		Meta: Meta{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
		},
		Errors: errorMessage,
	})
}
