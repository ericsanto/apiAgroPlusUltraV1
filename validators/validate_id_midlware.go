package validators

import "github.com/gin-gonic/gin"

func GetAndValidateIdMidlware(c *gin.Context, nameIdValidatedInMidlware string) uint {

	val, exists := c.Get(nameIdValidatedInMidlware)
	if !exists {
		return 0
	}

	id := val.(uint)

	return id
}
