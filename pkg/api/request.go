package api

import (
	"github.com/labstack/echo"
)

func GetBody(c echo.Context, req interface{}) error {
	err := c.Bind(req)

	if err == nil {
		err = c.Validate(req)
	}

	return err
}
