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
