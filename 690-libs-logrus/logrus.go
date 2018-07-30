package main

import (
	"math/rand"
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {

	log := logrus.New()
	// log.Formatter = &logrus.JSONFormatter{}

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

		log.Warning("aaa")

	}

}
