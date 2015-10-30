package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 1 * * * *", func() { fmt.Println("Every minute") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.AddFunc("@every 1s", func() { fmt.Println("Every one second") })
	c.Start()

	fmt.Println(c.Entries())
	for {
	}
}
