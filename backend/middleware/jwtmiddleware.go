package middleware

import (
	"net/http"
	"strings"

	"comb.com/banking/utils/jwt"
	echo "github.com/labstack/echo/v4"

	"fmt"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Println("Auth Header:", authHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "a"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": err})
		}

		// lưu thông tin user vào context
		c.Set("user", claims)
		return next(c)
	}
}
