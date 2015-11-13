package main

import (
	"fmt"
	"gopkg.in/validator.v2"
)

type NewUserRequest struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]$"`
	Name     string `validate:"nonzero"`
	Age      int    `validate:"min=21"`
	Password string `validate:"min=8"`
}

func main() {
	nur := NewUserRequest{Username: "something", Age: 20, Password: "aa"}

	if errs := validator.Validate(nur); errs != nil {
		fmt.Println(errs)
	}
}
