package handler

import "github.com/labstack/echo/v4"

type userHandler struct {
}

type UserHandler interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (uh *userHandler) Login(ctx echo.Context) error {
	// Implement login logic here
	return ctx.JSON(200, map[string]string{"message": "Logged in successfully"})
}

func (uh *userHandler) Register(ctx echo.Context) error {
	// Implement registration logic here
	return ctx.JSON(200, map[string]string{"message": "Registered successfully"})
}
