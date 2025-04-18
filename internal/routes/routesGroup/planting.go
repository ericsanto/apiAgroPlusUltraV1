package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterPlanting(r *gin.Engine) {

	plantingRepository := repositories.NewPlantingRepository(config.DB)
	plantingService := services.NewPlantingService(plantingRepository)
	plantingController := controllers.NewPlantingController(plantingService)

	plantingRouter := r.Group("/v1/plantacoes")
	plantingRouter.POST("/", plantingController.PostPlanting)
	plantingRouter.PUT("/:id", middlewares.ValidateIdParam("id"), plantingController.PutPlanting)
	plantingRouter.GET("/", middlewares.ValidateQueryParamPlanting(), plantingController.GetPlantingQueryParamBatchNameOrActive)
	plantingRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), plantingController.DeletePlanting)
}
