package router

import (
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(r *echo.Group) {
	r.GET("", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Welcome to User routes",
		})
	})
}
