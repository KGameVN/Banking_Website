package middleware

import (
	"github.com/labstack/echo/v4"
	"comb.com/banking/errors"
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