package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}) error {
	traceID := c.Response().Header().Get(echo.HeaderXRequestID)
	return c.JSON(http.StatusOK, APIResponse{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceID: traceID,
	})
}

func Created(c echo.Context, data interface{}) error {
	traceID := c.Response().Header().Get(echo.HeaderXRequestID)
	return c.JSON(http.StatusOK, APIResponse{
		Code:    1,
		Message: "Created",
		Data:    data,
		TraceID: traceID,
	})
}
