package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint8  `validate:"gte=0,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            13,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000000",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println("err1 ", err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("err2 ", err)
			return
		}
	}
}
