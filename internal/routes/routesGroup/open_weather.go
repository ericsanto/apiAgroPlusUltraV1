package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RouterGroupOpenWeather(r *gin.Engine) {

	openWeather := r.Group("/v1/weather-current")

	openWeather.GET("/", middlewares.GetCoordinates(), controllers.CurrentOpenWeather)
}
