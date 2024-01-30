package main

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string     //`validate:"required"`
	LastName  string     `validate:"required"`
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Addresses []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

type Address struct {
	Street  string `validate:"required"`
	City    string `validate:"required"`
	Country string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate
var ctx context.Context

func main() {
	fmt.Println("Run")
	validate = validator.New()
	address := &Address{
		Street:  "Cruz Ramirez 1",
		Country: "Spain",
		City:    "Madrit",
	}

	user := &User{
		FirstName: "",
		LastName:  "Rodriguez",
		Age:       35,
		Email:     "manuel.rodriguez@gmail.com",
		Addresses: []*Address{address},
	}

	err := validate.Struct(user)

	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		fmt.Println(errs)

		for _, err := range errs {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
		// return CreateClientResult{}, errors.Wrap(ErrValidationFailed, errs.Error())
	}

}
