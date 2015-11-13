package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

func main() {
	var validate *validator.Validate

	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Field(myEmail, "required,email")

	if errs != nil {
		fmt.Printf("%+v\n", errs)
		return
	}
}
