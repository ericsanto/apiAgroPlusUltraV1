package openweather

import (
	"fmt"
	"log"

	owm "github.com/briandowns/openweathermap"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type Main struct {
	Temperature    float64 `json:"temp"`
	TemperatureMax float64 `json:"temp_max"`
	TemperatureMin float64 `json:"temp_min"`
	FeelsLike      float64 `json:"feels_like"`
	Pressure       float64 `json:"pressure"`
	Humidity       int     `json:"humidity"`
}

type Rain struct {
	OneH   float64 `json:"1h"`
	ThreeH float64 `json:"3h"`
}

type Wind struct {
	Deg   float64 `json:"deg"`
	Speed float64 `json:"speed"`
}

type ResponseOpenWeather struct {
	Main     Main   `json:"main"`
	Rain     Rain   `json:"rain"`
	Wind     Wind   `json:"wind"`
	CityName string `json:"city"`
}

type OpenWeatherInterface interface {
	CurrentOpenWeather(latitude, longitude float64) (*ResponseOpenWeather, error)
	GetSolarRadiation(latitude, longitude float64) (float64, error)
}

type OpenWeather struct {
	apiKey string
}

func NewOpenWeather(apiKey string) OpenWeatherInterface {
	return &OpenWeather{apiKey: apiKey}
}

func (ow *OpenWeather) CurrentOpenWeather(latitude, longitude float64) (*ResponseOpenWeather, error) {

	w, err := owm.NewCurrent("C", "pt", ow.apiKey)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%w: %v", myerror.ErrNewCurrent, err)
	}

	err = w.CurrentByCoordinates(&owm.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%w: %v", myerror.ErrSearchCurrentByCoordinatesOpenWeather, err)
	}

	responseOpenWeather := ResponseOpenWeather{
		Main: Main{
			Temperature:    w.Main.Temp,
			TemperatureMax: w.Main.TempMax,
			TemperatureMin: w.Main.TempMin,
			FeelsLike:      w.Main.FeelsLike,
			Humidity:       w.Main.Humidity,
			Pressure:       w.Main.Pressure,
		},

		Rain: Rain{
			OneH:   w.Rain.OneH,
			ThreeH: w.Rain.ThreeH,
		},

		Wind: Wind{
			Deg:   w.Wind.Deg,
			Speed: w.Wind.Speed,
		},

		CityName: w.Name,
	}

	return &responseOpenWeather, nil

}

func (ow *OpenWeather) GetSolarRadiation(latitude, longitude float64) (float64, error) {

	uv, err := owm.NewUV(ow.apiKey)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("%w: %v", myerror.ErrGetUVSolarRadiationOpenWeather, err)
	}

	coord := &owm.Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
	}

	if err := uv.Current(coord); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("%w: %v", myerror.ErrGetUVSolarRadiationOpenWeather, err)
	}

	return uv.Value, nil

}
