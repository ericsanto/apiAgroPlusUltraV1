package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterIrrigationType(r *gin.Engine) {

	irrigationTypeController := di.NewIrrigationTypeBuilder().Builder()

	routerIrrigatinType := r.Group("/v1/tipo-irrigacao")

	routerIrrigatinType.GET("/", irrigationTypeController.GetAllIrrigationType)
	routerIrrigatinType.POST("/", irrigationTypeController.PostIrrigationType)
	routerIrrigatinType.GET("/:id", middlewares.ValidateIdParam("id"), irrigationTypeController.GetIrrigationTypeByID)
	routerIrrigatinType.PUT("/:id", middlewares.ValidateIdParam("id"), irrigationTypeController.PutIrrigationType)
	routerIrrigatinType.DELETE("/:id", middlewares.ValidateIdParam("id"), irrigationTypeController.DeleteIrrigationType)

}
