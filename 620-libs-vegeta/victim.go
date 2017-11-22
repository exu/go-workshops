package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}
