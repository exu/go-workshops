package main

import (
	"github.com/gen2brain/beeep"
)

// it should beep but i don't have beeper
func main() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}

	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}

	err = beeep.Notify("Title", "Message body")
	if err != nil {
		panic(err)
	}

}
