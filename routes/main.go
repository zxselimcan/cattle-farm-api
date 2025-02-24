package routes

import (
	"api/types"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) []*types.ErrorResponse {
	var validate = validator.New()
	var errors []*types.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element types.ErrorResponse
			element.Message = fmt.Sprintf("%v rules: %v: %v", err.Field(), err.Tag(), err.Param())
			errors = append(errors, &element)
		}
	}
	return errors
}
