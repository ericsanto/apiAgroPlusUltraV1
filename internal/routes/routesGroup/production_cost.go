package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupProductionCostRouter(r *gin.Engine) {

	productionCostRepository := repositories.NewProductionCostRepository(config.DB)
	productionCostService := services.NewProductionCostService(productionCostRepository)
	productionCostController := controllers.NewProductionCostController(productionCostService)

	productionCostRouter := r.Group("/v1/custos-plantacoes")
	productionCostRouter.GET("/", productionCostController.GetAllProductionCost)
	productionCostRouter.POST("/", productionCostController.PostProductionCost)
	productionCostRouter.GET("/:id", middlewares.ValidateIdParam("id"), productionCostController.GetProductionCostByID)
	productionCostRouter.PUT("/:id", middlewares.ValidateIdParam("id"), productionCostController.PutProductionCost)
	productionCostRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), productionCostController.DeleteProductionCost)

}
