package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

func SetupRouterIrrigationType(r *gin.Engine) {

	irrigationTypeRepository := repositories.NewIrrigationTypeRepository(config.DB)
	irrigationTypeService := services.NewIrrigationTypeService(irrigationTypeRepository)
	irrigatinTypeController := controllers.NewIrrigationTypeController(irrigationTypeService)

	routerIrrigatinType := r.Group("/v1/tipo-irrigacao")

	routerIrrigatinType.GET("/", irrigatinTypeController.GetAllIrrigationType)
	routerIrrigatinType.POST("/", irrigatinTypeController.PostIrrigationType)
	routerIrrigatinType.GET("/:id", middlewares.ValidateIdParam("id"), irrigatinTypeController.GetIrrigationTypeByID)
	routerIrrigatinType.PUT("/:id", middlewares.ValidateIdParam("id"), irrigatinTypeController.PutIrrigationType)
	routerIrrigatinType.DELETE("/:id", middlewares.ValidateIdParam("id"), irrigatinTypeController.DeleteIrrigationType)

}
