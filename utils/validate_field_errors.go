package utils

import (
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
)


func ValidateFieldErrors422UnprocessableEntity(entity interface{}) (map[string]string, error) {

  valid := validator.New()

  err := valid.Struct(entity)
  if err != nil {
    
    validatorErrors := err.(validator.ValidationErrors)
    errorMessages := make(map[string]string)

    for _, fieldErrors := range validatorErrors {
      errorMessages[strings.ToLower(fieldErrors.Field())] = fmt.Sprintf("%s: %s",fieldErrors, fieldErrors.Tag())  
    }

    return errorMessages, nil
  }

  return nil, err
}
  
