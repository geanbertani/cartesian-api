package api

import (
	"net/http"

	"github.com/geanbertani/cartesian-api/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoServer *echo.Echo

// Make server
func Make() *echo.Echo {
	echoServer = echo.New()

	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.Logger())

	// Initialize server
	echoServer.GET("/", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, echo.Map{common.MESSAGE: "Running server"})
	})

	return echoServer
}

// Instance task of echo
func Instance(task func(e *echo.Echo)) {
	task(echoServer)
}

// Run server
func Run() {
	// Start
	port := "9000"
	err := echoServer.Start(":" + port)

	// Log API errors
	echoServer.Logger.Fatal(err)
}
