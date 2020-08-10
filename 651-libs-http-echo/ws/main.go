package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

type Handler struct {
	Message string
}

func (handler *Handler) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		println("MiddleWare")
		return nil
	}
}

func (handler *Handler) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hej")
}

func (handler *Handler) WS(ws *websocket.Conn) {
	for {
		websocket.Message.Send(ws, "Hello, Client!"+handler.Message)
		msg := ""
		websocket.Message.Receive(ws, &msg)
		println(msg)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	handler := Handler{"Hoł hoł"}
	e.GET("/ws", echo.WrapHandler(websocket.Handler(handler.WS)))

	e.Start(":1323")
}
