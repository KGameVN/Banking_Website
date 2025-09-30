package converter

import (
	"strconv"
	"github.com/labstack/echo/v4"
)

func StringToInt(in string) (int, error) {
	number, err := strconv.Atoi(in)
	if err != nil {
		return -1, err
	}
	return number, nil
}

func StringToInt64(in string) (int64, error) {
	number, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return -1, err
	}
	return number, nil
}

func JsonToMap(c echo.Context) (map[string]interface{}, error) {
	var input map[string]interface{}
	if err := c.Bind(&input); err != nil {
		return nil, err
	}
	return input, nil
}
