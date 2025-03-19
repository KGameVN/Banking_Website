package utils

import (
	"github.com/labstack/echo"
)

func JsonToMap(c echo.Context) (map[string]interface{}, error) {
	var input map[string]interface{}
	if err := c.Bind(&input); err != nil {
		return nil, err
	}
	return input, nil
}
