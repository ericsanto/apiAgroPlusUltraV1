package middlewares

import (
	"net/http"
	"strconv"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
)

func ValidateQueryParamPestAgricultureCulture() gin.HandlerFunc {

	return func(c *gin.Context) {

		queryPestId := c.Query("pestId")
		queryCultureId := c.Query("cultureId")

		if queryPestId == "" || queryCultureId == "" {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "os parâmetros pestId e cultureId são obrigatórios",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		validatedPestId, err := strconv.ParseUint(queryPestId, 10, 32)
		if err != nil {
			return
		}

		validatedCultureId, err := strconv.ParseUint(queryCultureId, 10, 32)
		if err != nil {
			return
		}

		c.Set("validatedPestId", uint(validatedPestId))
		c.Set("validatedCultureId", uint(validatedCultureId))
		c.Next()
	}

}

func MiddlewareValidateQueryParamAgricultureCultureIdAndIrrigationRecomendedId() gin.HandlerFunc {
	return func(c *gin.Context) {

		queryCultureId := c.Query("cultureId")
		queryIrrigationRecomended := c.Query("irrigationId")

		if queryCultureId == "" || queryIrrigationRecomended == "" {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "os parâmetros cultureId e irrigationId são obrigatórios e não podem ser uma string vazia",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		validateCultureId, err := strconv.ParseUint(queryCultureId, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "o parâmtro cultureId deve ser um número inteiro não negativo",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		validateIrrigationId, err := strconv.ParseInt(queryIrrigationRecomended, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "o parâmtro irrigationId deve ser um número inteiro não negativo",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.Set("validatedCultureId", uint(validateCultureId))
		c.Set("validatedIrrigationId", uint(validateIrrigationId))
		c.Next()
	}
}
