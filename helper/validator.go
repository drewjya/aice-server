package helper

import (
	"log"

	"github.com/go-playground/validator"
)

func ValidateStruct(validate *validator.Validate, data interface{}) interface{} {
	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			log.Println(err)
		} else {
			var errorValue string
			for _, err := range err.(validator.ValidationErrors) {

				log.Println(err.Value())
				errorValue = errorValue + err.Value().(string) + "\n"

			}

			return errorValue
		}

		return err.Error()
	}
	return nil
}
