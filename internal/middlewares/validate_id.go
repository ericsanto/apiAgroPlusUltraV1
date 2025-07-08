package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

func ValidateIdParam(id string) gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Param(id)

		idUint, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, myerror.ErrorApp{
				Code:      http.StatusBadRequest,
				Message:   "O id dever ser um número inteiro",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.Set(id, uint(idUint))
		c.Next()
	}
}
