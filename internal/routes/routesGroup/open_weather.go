package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func RouterGroupOpenWeather(r *gin.Engine) {

	openWeatherController := di.NewOpenWeatherBuilder().Builder()

	openWeather := r.Group("/v1/weather-current")

	openWeather.GET("/", middlewares.GetCoordinates(), openWeatherController.CurrentOpenWeather)
}
