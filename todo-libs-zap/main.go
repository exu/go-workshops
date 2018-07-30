package main

//https://github.com/uber-go/zap

import (
	"time"

	"github.com/uber-go/zap"
)

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}
