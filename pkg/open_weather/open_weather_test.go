package openweather

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentOpenWeather_Success(t *testing.T) {

	mockOpenWeather := new(OpenWeatherMock)

	longitude := -10.245
	latitude := 69.904

	main := Main{
		Temperature:    20,
		TemperatureMax: 22,
		TemperatureMin: 15,
		FeelsLike:      22,
		Pressure:       10,
		Humidity:       12,
	}

	rain := Rain{
		OneH:   1,
		ThreeH: 0,
	}

	wind := Wind{
		Deg:   10.2,
		Speed: 20.2,
	}

	responseOpenWeather := ResponseOpenWeather{
		Main:     main,
		Rain:     rain,
		Wind:     wind,
		CityName: "Lagarto",
	}

	mockOpenWeather.On("CurrentOpenWeather", latitude, longitude).Return(&responseOpenWeather, nil)

	resp, err := mockOpenWeather.CurrentOpenWeather(latitude, longitude)

	assert.Nil(t, err)
	assert.Equal(t, responseOpenWeather.Main.Temperature, resp.Main.Temperature)
	assert.EqualValues(t, responseOpenWeather.CityName, "Lagarto")

	mockOpenWeather.AssertExpectations(t)
}

func TestCurrentOpenWeather_Error(t *testing.T) {

	mockOpenWeather := new(OpenWeatherMock)

	longitude := -10.245
	latitude := 69.904

	mockOpenWeather.On("CurrentOpenWeather", latitude, longitude).Return(&ResponseOpenWeather{}, fmt.Errorf("erro ao buscar dados climaticos a partir das corrdenadas fornecidas"))

	resp, err := mockOpenWeather.CurrentOpenWeather(latitude, longitude)

	assert.NotNil(t, err)
	assert.EqualValues(t, resp.Main.Temperature, 0)
	assert.EqualError(t, err, "erro ao buscar dados climaticos a partir das corrdenadas fornecidas")
	assert.EqualValues(t, resp.CityName, "")

	mockOpenWeather.AssertExpectations(t)

}

func TestGetSolarRadiation(t *testing.T) {

	mockRepo := new(OpenWeatherMock)

	latitude := -50.38
	longitude := 32.90

	solarRadiation := 8.2

	mockRepo.On("GetSolarRadiation", latitude, longitude).Return(solarRadiation, nil)

	respSolarRadiation, err := mockRepo.GetSolarRadiation(latitude, longitude)

	assert.Nil(t, err)
	assert.Equal(t, solarRadiation, respSolarRadiation)
	assert.IsType(t, solarRadiation, respSolarRadiation)

	mockRepo.AssertExpectations(t)
}
