package middlewares

import "github.com/gin-gonic/gin"

func ValidateQueryParamPlanting() gin.HandlerFunc {

	return func(c *gin.Context) {

		batchName := c.Query("batchName")
		active := c.Query("active")

		c.Set("batchName", batchName)
		c.Set("active", active)
		c.Next()

	}
}
