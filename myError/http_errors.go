package myerror

import (
	"time"

	"github.com/gin-gonic/gin"
)

// if err := c.ShouldBindJSON(&requestAgricultureCultureIrrigation); err != nil {
// 	c.JSON(http.StatusBadRequest, myerror.ErrorApp{
// 		Code:      http.StatusBadRequest,
// 		Message:   "body da requisição é inválido",
// 		Timestamp: time.Now().Format(time.RFC3339),
// 	})
// 	return
// }

func HttpErrors(statusCode int, message interface{}, c *gin.Context) {
	c.AbortWithStatusJSON(statusCode, ErrorApp{
		Code:      statusCode,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
