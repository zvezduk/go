package api

import (
	"github.com/labstack/echo"
)

func Public(e *echo.Echo) {
	// Базовые ресурсы
	e.GET("/resources", getResources)
	e.POST("/resources", setResources)
}
