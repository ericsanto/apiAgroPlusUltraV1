package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
)

type OpenWeatherControllerInterface interface {
	CurrentOpenWeather(c *gin.Context)
}

type OpenWeatherController struct {
	openWeatherService openweather.OpenWeatherInterface
}

func NewOpenWeather(openWeatherService openweather.OpenWeatherInterface) OpenWeatherControllerInterface {
	return &OpenWeatherController{openWeatherService: openWeatherService}
}

func (opwc *OpenWeatherController) CurrentOpenWeather(c *gin.Context) {

	val, _ := c.Get("lat")

	lat := val.(float64)

	val, _ = c.Get("long")

	long := val.(float64)

	response, err := opwc.openWeatherService.CurrentOpenWeather(lat, long)

	if err != nil {
		myerror.HttpErrors(http.StatusBadGateway, "problema de comunicacao com servidor externo", c)
		return
	}

	c.JSON(http.StatusOK, response)

}
