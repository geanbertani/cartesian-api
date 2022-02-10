package points

import (
	"github.com/geanbertani/cartesian-api/libs/handler"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	auth := e.Group("api")
	auth.GET("/points", GetPointsByDistanceHandler, handler.MiddlewareBindAndValidate(&Filter{}))
}
