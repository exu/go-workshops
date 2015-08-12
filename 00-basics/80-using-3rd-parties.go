package main

// wystarczy w katalogu projektu użyć komendy `go get`
// paczka zostanie sciagnieta i umieszczona w katalogach src, bin i pkg
// przy standardowej konfirguracji a-la global

import (
	"math/rand"

	"github.com/labstack/echo"
)

func hello(c *echo.Context) error {
	// możesz użyc stałych z pakietu  net/http np: http.StatusOK
	// najpierw import trzeba zrobić (import net/http)
	return c.JSON(200, map[string]int{
		"nmel_errors_count": rand.Int(),
		"users":             rand.Int(),
	})
}

func main() {
	e := echo.New()
	e.Get("/", hello)
	e.Run(":8080")
}
