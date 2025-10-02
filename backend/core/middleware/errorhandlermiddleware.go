package middleware

import (
	"comb.com/banking/internal/errors"
	"github.com/labstack/echo/v4"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			status, payload := errors.HTTPStatusAndPayload(err)
			return c.JSON(status, payload)
		}
		return nil
	}
}
