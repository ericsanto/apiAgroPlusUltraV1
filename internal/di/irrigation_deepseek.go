package di

import (
	"fmt"
	"os"

	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/deepseek"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/mosquitto"
	openweather "github.com/ericsanto/apiAgroPlusUltraV1/pkg/open_weather"
)

type IrrigationDeepseekServiceBuilder struct{}

func NewIrrigationDeepSeekBuilder() *IrrigationDeepseekServiceBuilder {
	return &IrrigationDeepseekServiceBuilder{}
}

func (idsb *IrrigationDeepseekServiceBuilder) Builder() (controllers.IrrigationRecommendedDeepSeekControllerInterface, error) {

	openWeather := openweather.NewOpenWeather(os.Getenv("OPEN_WEATHER_API_KEY"))

	opts := mosquitto.OPTSMosquitto{
		UrlBrokerMosquitto: os.Getenv("URL_BROKER_MOSQUITTO"),
		Username:           os.Getenv("USERNAME_BROKER_MOSQUITTO"),
		Password:           os.Getenv("PASSWORD_BROKER_MOSQUITTO"),
		CleanSession:       false,
		ClientID:           "golang-api",
	}

	jsonUtils := jsonutil.NewJsonUtils()
	mosquittoClient, err := mosquitto.NewMosquittoBroker(opts, jsonUtils)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar build do IrrigationDeepseek %w", err)
	}

	deepseekApiKey := os.Getenv("DEEPSEEK_API_KEY")

	deepseekClient := deepseek.NewDeepSeek(deepseekApiKey)

	farmRepository := repositories.NewFarmRepository(db.DB)

	plantingRepository := repositories.NewPlantingRepository(db.DB, farmRepository)
	irrigationDeepseekService := services.NewIrrigationRecomendedDeepseekService(plantingRepository, openWeather, mosquittoClient, deepseekClient)
	irrigationDeeepeekControler := controllers.NewIrrigationRecommendedDeepseekController(irrigationDeepseekService)

	return irrigationDeeepeekControler, nil
}
