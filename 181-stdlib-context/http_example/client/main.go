package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// creating context with timeout
	ctx := context.Background()

	// We also can add default timeout to context
	//ctx, cancel := context.WithTimeout(ctx, time.Second)
	// defer cancel to clear context (in case if it will not be called in timeout)
	//defer cancel()

	// creating request with context
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	// do request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout, res.Body)

	defer res.Body.Close()
}
