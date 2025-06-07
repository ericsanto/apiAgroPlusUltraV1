package validators

import (
	"fmt"
	"net/http"
	"strings"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateFieldErrors422UnprocessableEntity(entity interface{}) (map[string]string, error) {

	valid := validator.New()

	err := valid.Struct(entity)
	if err != nil {

		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)

		for _, fieldErrors := range validatorErrors {
			errorMessages[strings.ToLower(fieldErrors.Field())] = fmt.Sprintf("%s: %s", fieldErrors, fieldErrors.Tag())
		}

		return errorMessages, nil
	}

	return nil, err
}

func ValidateShouldBindJson(structRequest interface{}, c *gin.Context) interface{} {

	if err := c.ShouldBind(&structRequest); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
	}

	return nil
}

// func ValidateRequest(structRequest interface{}, c *gin.Context) interface{}  {

// 	validate, err := ValidateFieldErrors422UnprocessableEntity(structRequest)
// 	if err != nil {
// 		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
// 	}

// 	if len(validate) > 0 {
// 		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)

// 	}

// }

// var requestBatch requests.BatchRequest

// 	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestBatch)
// 	if err != nil {
// 		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
// 		return
// 	}

// 	if len(validate) > 0 {
// 		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
// 		return
// 	}

// 	validatedID := validators.GetAndValidateIdMidlware(c, "validatedID")

// 	if err := b.batchService.PutBatch(validatedID, requestBatch); err != nil {
// 		if strings.Contains(err.Error(), "já existe lote cadastrado com esse nome") {
// 			myerror.HttpErrors(http.StatusConflict, "já existe lote cadastrado com esse nome", c)
// 			return
// 		}

// 		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
// 		return
// 	}

// 	c.Status(http.StatusOK)
