package openweather

import (
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")

type Main struct {
	Temperature    float64 `json:"temp"`
	TemperatureMax float64 `json:"temp_max"`
	TemperatureMin float64 `json:"temperature_min"`
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

type ResponseOpenWheather struct {
	Main     Main   `json:"main"`
	Rain     Rain   `json:"rain"`
	Wind     Wind   `json:"wind"`
	CityName string `json:"city"`
}

func CurrentOpenWeather(latitude, longitude float64) (interface{}, error) {

	w, err := owm.NewCurrent("C", "pt", apiKey)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = w.CurrentByCoordinates(&owm.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	responseOpenWeather := ResponseOpenWheather{
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

	return responseOpenWeather, nil

}
