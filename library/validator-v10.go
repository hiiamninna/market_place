package library

import (
	"fmt"
	"strings"

	validation "github.com/go-playground/validator/v10"
)

// Using validator
func ValidateInput(data interface{}) (string, error) {

	// create new validationa and check the struct
	var validationError = validation.New()
	err := validationError.Struct(data)

	if err != nil {
		var errors []string
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validation.InvalidValidationError); ok {
			fmt.Println(err)
			return "", nil
		}

		for n, e := range err.(validation.ValidationErrors) {
			message := fmt.Sprintf("%s must %s,", e.Field(), e.Tag())
			if n == (len(err.(validation.ValidationErrors)) - 1) {
				if e.Tag() == "email" {
					message = "Please input correct email format"
				} else {
					message = fmt.Sprintf("%s must %s", e.Field(), e.Tag())
				}
			}
			errors = append(errors, message)
		}
		return fmt.Sprint(errors), err
	}

	if err != nil {
		arrayOfErrors := []string{err.Error()}
		return fmt.Sprint(arrayOfErrors), err
	}

	return "", err
}

func IsAnImageUrl(value string) bool {
	if strings.HasSuffix(strings.ToUpper(value), ".JPG") || strings.HasSuffix(strings.ToUpper(value), ".JPEG") {
		return true
	}
	return false
}
