package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ValidateQueryParamAgricultureCulturePestMethod() gin.HandlerFunc {
	return func(c *gin.Context) {

		agricultureCulureName := c.Query("agricultureCultureName")
		pestName := c.Query("pestName")
		sustainablePestControlMethod := c.Query("sustainablePestControlMethod")

		c.Set("agricultureCultureName", agricultureCulureName)
		c.Set("pestName", pestName)
		c.Set("sustainablePestControlMethod", sustainablePestControlMethod)
		c.Next()
	}
}
