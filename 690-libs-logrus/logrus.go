package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/weekface/mgorus"
	"math/rand"
	"time"
)

func main() {

	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	hooker, err := mgorus.NewHooker("localhost:27017", "kasia", "logs")
	if err == nil {
		log.Hooks.Add(hooker)
	}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	type Complex struct {
		ID          string
		Name        string
		From        time.Time
		To          time.Time
		Description string
		Count       int
	}

	for {
		log.WithFields(logrus.Fields{
			"number": rand.Int(),
			"item1":  &Complex{"aaa", "Something", time.Now(), time.Now(), "long description", 3},
			"item2":  &Complex{"aaa", "Else", time.Now(), time.Now(), "long description", 3},
		}).Error("Error appears")

	}

}
