package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		fmt.Println(err)
		// TODO: handling many type error
		// Optionally, you could return the error to give each route more control over the status code
		return ErrBadRequest
	}
	return nil
}
