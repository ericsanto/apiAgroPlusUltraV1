package services

import (
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/deepseek"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/mosquitto"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
)

type IrrigationRecommendedDeepSeekServiceInterface interface {
	IrrigationRecommendedDeepSeek(latitude, longitude float64, userID, farmID uint) error
}

type IrrigationRecommendedDeepSeekService struct {
	plantingRepository repositories.PlantingRepositoryInterface
	openWeather        openweather.OpenWeatherInterface
	mosquittoClient    mosquitto.MosquittoInterface
	deepseekClient     deepseek.LLMMethods
}

func NewIrrigationRecomendedDeepseekService(plantingRepository repositories.PlantingRepositoryInterface, openWeather openweather.OpenWeatherInterface,
	mosquittoClient mosquitto.MosquittoInterface, deepseekClient deepseek.LLMMethods) IrrigationRecommendedDeepSeekServiceInterface {
	return &IrrigationRecommendedDeepSeekService{plantingRepository: plantingRepository,
		openWeather:     openWeather,
		mosquittoClient: mosquittoClient,
		deepseekClient:  deepseekClient}
}

func (p *IrrigationRecommendedDeepSeekService) IrrigationRecommendedDeepSeek(latitude, longitude float64, userID, farmID uint) error {

	responseOpenWeather, err := p.openWeather.CurrentOpenWeather(latitude, longitude)

	if err != nil {
		log.Println(err)
		return err
	}

	isPlanting := true

	plantingIsTrue, err := p.plantingRepository.FindByParamBatchNameOrIsActivePlanting("", isPlanting, userID, farmID)

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

	solarRadiation, err := p.openWeather.GetSolarRadiation(latitude, longitude)
	if err != nil {
		return err
	}

	for _, v := range plantingIsTrue {

		valuesPromptModel := deepseek.PromptModel{
			CurrentTemperature:         temperature,
			RelativeAirHumidity:        float64(humidityAirRelative),
			WindSpeed:                  windSpeed,
			WindDirection:              windDirection,
			SoilHumidity:               float64(soilHumidity),
			SystemIrrigationEfficiency: float64(systemIrrigationEfficiency),
			SpaceBetweenPlants:         v.SpaceBetweenPlants,
			SpaceBetweenRows:           v.SpaceBetweenRows,
			AtmosphericPressure:        atmosphericPressure,
			StartDatePlanting:          v.StartDatePlanting.String(),
			SoilType:                   v.SoilTypeName,
			IrrigationType:             v.IrrigationType,
			AgricultureCulture:         v.AgricultureCultureName,
			BatchName:                  v.BatchName,
			SolarRadiation:             solarRadiation,
		}

		content, err := p.deepseekClient.RequestRecommendationIrrigation(valuesPromptModel)

		if err != nil {
			log.Println(err.Error())
			return err
		}

		jsons, err := p.mosquittoClient.FormatResponseDeepSeekInJSONForMqttBrokerPublisher(content, "```json", "```")

		if err != nil {
			log.Println(err.Error())
			return err
		}

		isSendMessage, err := p.mosquittoClient.Publisher("irrigation", jsons)

		if !isSendMessage {
			log.Println(err.Error())
			continue
		}
	}

	p.mosquittoClient.Disconnect(250)

	return nil
}
