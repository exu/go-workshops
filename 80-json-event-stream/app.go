package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)

func apiHandler(c *echo.Context) error {
	i := 0
	c.Response().Header().Set(echo.ContentType, "text/event-stream")
	c.Response().WriteHeader(http.StatusOK)

	for {
		l := locations[i]
		c.Response().Write([]byte("data: "))
		if err := json.NewEncoder(c.Response()).Encode(l); err != nil {
			return err
		}
		c.Response().Write([]byte("\n\n"))
		c.Response().Flush()
		time.Sleep(1 * time.Second)

		i++
		if i >= len(locations) {
			return nil
			i = 0
		}

	}
	return nil
}

func main() {
	e := echo.New()
	e.Static("/", "./")
	e.Get("/api", apiHandler)
	e.Run(":6788")
}
