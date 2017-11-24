package main

import (
	"fmt"
	"net/http"

	"github.com/bamzi/jobrunner"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	jobrunner.Start(10, 1)
	jobrunner.Schedule("@every 5s", ReminderEmails{})
	jobrunner.Schedule("@every 1s", SendMeSomething{})

	// Połączymy scheduler z serwisem WWW
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.GET("/status", status)
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func status(c echo.Context) error {
	return c.JSON(http.StatusOK, jobrunner.StatusJson())
}

type ReminderEmails struct {
}

func (e ReminderEmails) Run() {
	fmt.Printf("Every 5 sec send reminder emails \n")
}

type SendMeSomething struct {
}

func (e SendMeSomething) Run() {
	fmt.Printf("Every 1 sec send me something \n")
}
