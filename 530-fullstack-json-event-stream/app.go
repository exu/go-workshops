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

func apiHandler(c echo.Context) error {
	i := 0
	response := c.Response()

	// ustawiamy header na event-stream rozpoznawalny przez przeglądarkę
	response.Header().Set(echo.HeaderContentType, "text/event-stream")
	response.WriteHeader(http.StatusOK) // 200

	for {
		location := locations[i]
		response.Write([]byte("data: "))

		if err := json.NewEncoder(response).Encode(location); err != nil {
			return err
		}

		response.Write([]byte("\n\n"))
		response.Flush()
		time.Sleep(100 * time.Millisecond)

		i++
		if i >= len(locations) {
			i = 0
		}

	}

	return nil
}

func main() {
	e := echo.New()
	e.Static("/", "./static")
	e.GET("/api", apiHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
