package middlewares

import (
	"net/http"
	"strconv"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
)

func MiddlewareValidateQueryParamAgricultureCultureIrrigation() gin.HandlerFunc {

	return func(c *gin.Context) {

		queryCultureId := c.Query("cultureId")

		if queryCultureId == "" {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "o parâmetro cultureId é obrigatório",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		validatedCultureId, err := strconv.ParseUint(queryCultureId, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "o parâmetro dever ser um número inteiro não negativo",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		// typeQueryCultureId := reflect.ValueOf(&validatedCultureId)
		// fmt.Println(typeQueryCultureId.Elem().Kind())

		// if typeQueryCultureId.Elem().Kind() != reflect.Uint64 {
		// 	c.JSON(http.StatusBadRequest, myerror.ErrorApp{
		// 		Code:      http.StatusBadRequest,
		// 		Message:   "o parâmetro dever ser um número inteiro não negativo",
		// 		Timestamp: time.Now().Format(time.RFC3339),
		// 	})
		// 	return
		// }

		c.Set("validatedCultureId", uint(validatedCultureId))
		c.Next()
	}

}
