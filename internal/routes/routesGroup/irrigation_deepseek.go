package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

func SetupRouterIrrigationDeepSeek(r *gin.Engine) {

	plantingRepository := repositories.NewPlantingRepository(config.DB)
	irrigationDeepseekService := services.NewIrrigationRecomendedDeepseekService(plantingRepository)
	irrigationDeeepeekControler := controllers.NewIrrigationRecommendedDeepseekController(irrigationDeepseekService)

	routerIrrigationDeepseek := r.Group("/v1/irrigation-deepseek")

	routerIrrigationDeepseek.GET("/", middlewares.GetCoordinates(), irrigationDeeepeekControler.IrrigationDeepseek)

}
