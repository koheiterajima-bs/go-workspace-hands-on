package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", helloHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
