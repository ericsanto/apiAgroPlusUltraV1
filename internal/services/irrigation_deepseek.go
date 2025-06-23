package services

import (
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/deepseek"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/mosquitto"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
)

type IrrigationRecommendedDeepSeekService struct {
	plantingRepository *repositories.PlantingRepository
}

func NewIrrigationRecomendedDeepseekService(plantingRepository *repositories.PlantingRepository) *IrrigationRecommendedDeepSeekService {
	return &IrrigationRecommendedDeepSeekService{plantingRepository: plantingRepository}
}

func (p *IrrigationRecommendedDeepSeekService) IrrigationRecommendedDeepSeek(latitude, longitude float64) error {

	responseOpenWeather, err := openweather.CurrentOpenWeather(latitude, longitude)

	if err != nil {
		log.Println(err)
		return err
	}

	clientMQTT, err := mosquitto.CreateClient()

	if err != nil {
		return err
	}

	isPlanting := true

	plantingIsTrue, err := p.plantingRepository.FindByParamBatchNameOrIsActivePlanting("", isPlanting)

	if err != nil {
		return err
	}

	temperature := responseOpenWeather.Main.Temperature
	humidityAirRelative := responseOpenWeather.Main.Humidity
	windSpeed := responseOpenWeather.Wind.Speed
	windDirection := responseOpenWeather.Wind.Deg
	soilHumidity := 90
	systemIrrigationEfficiency := 80
	atmosphericPressure := responseOpenWeather.Main.Pressure

	solarRadiation, err := openweather.GetSolarRadiation(latitude, longitude)
	if err != nil {
		return err
	}

	for _, v := range plantingIsTrue {

		content, err := deepseek.RequestRecommendationIrrigationDeepSeek(temperature, float64(humidityAirRelative), windSpeed,
			windDirection, solarRadiation, float64(soilHumidity), float64(systemIrrigationEfficiency),
			v.SpaceBetweenPlants, v.SpaceBetweenRows, atmosphericPressure, v.StartDatePlanting.String(),
			v.SoilTypeName, v.IrrigationType, v.AgricultureCultureName, v.BatchName)

		if err != nil {
			log.Println(err.Error())
			return err
		}

		jsons, err := mosquitto.FormatResponseDeepSeekInJSONForMqttBrokerPublisher(content, "```json", "```")

		if err != nil {
			log.Println(err.Error())
			return err
		}

		isSendMessage, err := mosquitto.Publisher(clientMQTT, "irrigation", jsons)

		if !isSendMessage {
			log.Println(err.Error())
			continue
		}
	}

	clientMQTT.Disconnect(250)

	return nil
}
