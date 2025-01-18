package main

import (
	"car-lending-api/middleware"
	"car-lending-api/router"

	"github.com/labstack/echo/v4"
)

var secretKey = []byte("secret-key")

func main() {
	e := echo.New()

	// Create JWT middleware
	jwtMiddleware := middleware.NewJWTMiddleware(secretKey)

	// Register User routes
	router.RegisterUserRoutes(e.Group("/user", jwtMiddleware))

	// Register Admin routes
	router.RegisterAdminRoutes(e.Group("/admin", jwtMiddleware))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
