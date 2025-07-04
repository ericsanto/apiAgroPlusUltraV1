package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupProductionCostRouter(r *gin.Engine) {

	productionCostController := di.NewProductionCostBuilder().Builder()

	productionCostRouter := r.Group("/v1/custos-plantacoes")
	productionCostRouter.GET("/", productionCostController.GetAllProductionCost)
	productionCostRouter.POST("/", productionCostController.PostProductionCost)
	productionCostRouter.GET("/:id", middlewares.ValidateIdParam("id"), productionCostController.GetProductionCostByID)
	productionCostRouter.PUT("/:id", middlewares.ValidateIdParam("id"), productionCostController.PutProductionCost)
	productionCostRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), productionCostController.DeleteProductionCost)

}
