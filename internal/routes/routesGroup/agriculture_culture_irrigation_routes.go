package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterAgricultureCultureIrrigation(r *gin.Engine) {

	agricultureCultureIrrigationRepository := repositories.NewAgricultureCultureIrrigationRepository(config.DB)
	agricultureCultureIrrigationService := services.NewAgricultureCultureIrrigationService(agricultureCultureIrrigationRepository)
	agricultureCultureIrrigationController := controllers.NewAgricultureCultureIrrigationController(agricultureCultureIrrigationService)

	agricultureCultureIrrigationRouter := r.Group("/v1/irrigacao-cultura")
	agricultureCultureIrrigationRouter.GET("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIrrigation(), agricultureCultureIrrigationController.GetAgricultureCultureIrrigationFindByIDController)
	agricultureCultureIrrigationRouter.POST("/", agricultureCultureIrrigationController.PostAgricultureCultureIrrigationController)
	agricultureCultureIrrigationRouter.DELETE("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIdAndIrrigationRecomendedId(), agricultureCultureIrrigationController.DeleteAgricultureCulturueIrrigation)
	agricultureCultureIrrigationRouter.PUT("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIdAndIrrigationRecomendedId(), agricultureCultureIrrigationController.PutAgricultureCultureIrrigation)
}
