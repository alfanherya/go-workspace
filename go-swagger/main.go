package main

import (
	"net/http"

	_ "github.com/alfanherya/go-workspace/go-swagger/docs/go-echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route
	e.GET("/", Healtcheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start server
	e.Logger.Fatal(e.Start(":4000"))
}

func Healtcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
