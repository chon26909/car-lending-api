package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// NewJWTMiddleware สร้าง middleware สำหรับตรวจสอบ JWT
func NewJWTMiddleware(secretKey []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ดึงและตรวจสอบ token
			tokenString, err := extractToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": err.Error(),
				})
			}

			// ตรวจสอบความถูกต้องของ token
			if err := verifyToken(tokenString, secretKey); err != nil {
				fmt.Println("error verifying token:", err)
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized: invalid token",
				})
			}

			return next(c)
		}
	}
}

// extractToken ดึง token จาก Header
func extractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("unauthorized: missing token")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("unauthorized: invalid token format")
	}

	return parts[1], nil
}

// verifyToken ตรวจสอบความถูกต้องของ token
func verifyToken(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
