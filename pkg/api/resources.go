package api

import (
	"github.com/labstack/echo"
	"github.com/zvezduk/go/pkg/resource"
	"net/http"
)

func getResources(c echo.Context) error {
	r := resource.GetList()

	return c.JSON(http.StatusOK, r)
}

func setResources(c echo.Context) error {
	r := new(resource.Resources)

	if err := GetBody(c, r); err != nil {
		return err
	}

	if err := resource.SetList(r); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
