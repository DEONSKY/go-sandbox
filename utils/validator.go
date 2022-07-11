package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(object interface{}) []string {
	var errors []string
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := "Violated ; Field: " + err.Field() + " , Rule : " + err.Tag() + " , " + err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}
