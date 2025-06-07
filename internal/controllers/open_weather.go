package controllers

import (
	"net/http"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
	"github.com/gin-gonic/gin"
)

func CurrentOpenWeather(c *gin.Context) {

	val, _ := c.Get("lat")

	lat := val.(float64)

	val, _ = c.Get("long")

	long := val.(float64)

	response, err := openweather.CurrentOpenWeather(lat, long)

	if err != nil {
		myerror.HttpErrors(http.StatusBadGateway, "problema de comunicacao com servidor externo", c)
		return
	}

	c.JSON(http.StatusOK, response)

}
