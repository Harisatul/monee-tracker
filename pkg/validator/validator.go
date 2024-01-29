package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validation(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" && e.Field() == "Email" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "email" {
					message = "Email yang anda masukkan tidak valid"
				}
			}
			return errors.New(message)
		}
	}
	return nil
}
