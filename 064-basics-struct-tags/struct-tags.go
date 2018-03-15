package main

import (
	"fmt"
	"reflect"
)

// lets define interface for our messanger which
// can return some message which will be displayed in terminal
type Messanger interface {
	GetMessage() string
}

// lets implement Error message with prefix and postfix tag
// where we can put some shell colors
type ErrorMessage struct {
	Message string `prefix:"\x1b[31m" postfix:"\x1b[0m"`
}

func (message ErrorMessage) GetMessage() string {
	return "ERROR: Booo! " + message.Message
}

// lets implement some LogMessage
// with same prefix and postfix tags
type LogMessage struct {
	Message string `prefix:"\x1b[32m" postfix:"\x1b[0m"`
}

func (message LogMessage) GetMessage() string {
	return message.Message
}

// function show will use our prefix and posfix tag and display message
// between those two.
func show(input Messanger) {
	st := reflect.TypeOf(input)
	field := st.Field(0)

	fmt.Printf(
		"%s%s%s\n",
		field.Tag.Get("prefix"),
		input.GetMessage(),
		field.Tag.Get("postfix"),
	)
}

// now show some logs and errors
func main() {
	show(LogMessage{"Hoł hoł"})
	show(ErrorMessage{Message: "Noł noł noł"})
}
