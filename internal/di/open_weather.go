package di

import (
	"os"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
)

type OpenWeatherBuilder struct{}

func NewOpenWeatherBuilder() *OpenWeatherBuilder {
	return &OpenWeatherBuilder{}
}

func (opw *OpenWeatherBuilder) Builder() controllers.OpenWeatherControllerInterface {

	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")

	openWeatherService := openweather.NewOpenWeather(openWeatherApiKey)

	openWeatherController := controllers.NewOpenWeather(openWeatherService)

	return openWeatherController
}
