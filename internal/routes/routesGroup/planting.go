package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterPlanting(r *gin.Engine) {

	plantingController := di.NewPlantingRepository().Builder()

	plantingRouter := r.Group("/v1/plantacoes")
	plantingRouter.POST("/", plantingController.PostPlanting)
	plantingRouter.PUT("/:id", middlewares.ValidateIdParam("id"), plantingController.PutPlanting)
	plantingRouter.GET("/", middlewares.ValidateQueryParamPlanting(), plantingController.GetPlantingQueryParamBatchNameOrActive)
	plantingRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), plantingController.DeletePlanting)
}
