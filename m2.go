package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

func main() {
	v := validator.New()
	a := User{
		Email: "a",
		Name:  "",
	}
	fmt.Println(a)
	err := v.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}
