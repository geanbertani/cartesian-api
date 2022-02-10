package points

import (
	"net/http"

	"github.com/eucatur/go-toolbox/handler"
	"github.com/geanbertani/cartesian-api/common"
	"github.com/labstack/echo/v4"
)

// FindByFilterHandler ...
func GetPointsByDistanceHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Filter)

	pointsDistance, err := GetPointsByDistance(p)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{common.MESSAGE: err.Error()})
	}

	return c.JSON(http.StatusOK, pointsDistance)
}
