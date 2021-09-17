package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/rizalgowandy/go-swag-sample/docs/echosimple"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routesgit p
	e.GET("/", Healtcheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start Server
	e.Logger.Fatal(e.Start(":4000"))
}

func Healtcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
