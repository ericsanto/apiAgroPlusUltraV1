package openweather

import "github.com/stretchr/testify/mock"

type OpenWeatherMock struct {
	mock.Mock
}

func (owm *OpenWeatherMock) CurrentOpenWeather(latitude, longitude float64) (*ResponseOpenWeather, error) {

	args := owm.Called(latitude, longitude)

	return args.Get(0).(*ResponseOpenWeather), args.Error(1)
}

func (owm *OpenWeatherMock) GetSolarRadiation(latitude, longitude float64) (float64, error) {

	args := owm.Called(latitude, longitude)

	return args.Get(0).(float64), args.Error(1)
}
