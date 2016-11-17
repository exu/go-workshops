package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
	"net/http"
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
	e.GET("/ws", standard.WrapHandler(websocket.Handler(handler.WS)))

	e.Run(standard.New(":1323"))
}
