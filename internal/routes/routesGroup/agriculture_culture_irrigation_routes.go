package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterAgricultureCultureIrrigation(r *gin.Engine) {

	agricultureCultureIrrigationController := di.NewAgricultureCultureIrrigationBuilder().Builder()

	agricultureCultureIrrigationRouter := r.Group("/v1/irrigacao-cultura")
	agricultureCultureIrrigationRouter.GET("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIrrigation(), agricultureCultureIrrigationController.GetAgricultureCultureIrrigationFindByIDController)
	agricultureCultureIrrigationRouter.POST("/", agricultureCultureIrrigationController.PostAgricultureCultureIrrigationController)
	agricultureCultureIrrigationRouter.DELETE("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIdAndIrrigationRecomendedId(), agricultureCultureIrrigationController.DeleteAgricultureCulturueIrrigation)
	agricultureCultureIrrigationRouter.PUT("/", middlewares.MiddlewareValidateQueryParamAgricultureCultureIdAndIrrigationRecomendedId(), agricultureCultureIrrigationController.PutAgricultureCultureIrrigation)
}
