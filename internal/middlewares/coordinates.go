package middlewares

import (
	"net/http"
	"strconv"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
)

func GetCoordinates() gin.HandlerFunc {

	return func(c *gin.Context) {

		latitude := c.Query("lat")
		longitude := c.Query("long")

		if latitude == "" || longitude == "" {
			myerror.HttpErrors(http.StatusBadRequest, "os parametros lat e long sao obrigatorios", c)
			return
		}

		validatedLat, err := strconv.ParseFloat(latitude, 64)
		if err != nil {
			return
		}

		validatedLong, err := strconv.ParseFloat(longitude, 64)
		if err != nil {
			return
		}

		c.Set("lat", float64(validatedLat))
		c.Set("long", float64(validatedLong))
		c.Next()
	}
}
