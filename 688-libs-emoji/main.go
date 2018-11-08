package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {

	// there is need to install noto-fonts-emoji
	// looks like working only on gnome-terminal in linux OutOfTheBox
	// Konsole and URXVT doesn't work at all.

	fmt.Println("Hello World Emoji!")
	emoji.Println(":beer: Beer!!!")
	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)
}
