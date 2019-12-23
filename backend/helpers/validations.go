package helpers

import (
	"github.com/cryptosalamander/gorm_crud_example/dtos"
	"github.com/cryptosalamander/gorm_crud_example/langs"
	"gopkg.in/go-playground/validator.v8"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var validations []dtos.Validation

	// get validation errors
	validationError := err.(validator.ValidationErrors)

	for _, value := range validationError {
		// get field & rule (tag)
		field, rule := value.Field(), value.Tag()
		// create validation object
		validation := dtos.Validation{Field: field, Message: langs.GenerateValidationMessage(field, rule)}

		// add validation object to validations
		validations = append(validations, validation)
	}

	// set Validations response
	response.Validations = validations

	return response
}
