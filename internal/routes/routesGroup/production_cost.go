package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupProductionCostRouter(r *gin.Engine) {

	productionCostController := di.NewProductionCostBuilder().Builder()

	productionCostRouter := r.Group("/v1/fazenda")

	productionCostRouter.GET("/:farmID/lote/:batchID/plantacoes/:plantingID/custos",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		productionCostController.GetAllProductionCost)

	productionCostRouter.POST("/:farmID/lote/:batchID/plantacoes/:plantingID/custos",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		productionCostController.PostProductionCost)

	productionCostRouter.GET("/:farmID/lote/:batchID/plantacoes/:plantingID/custos/:costID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("costID"),
		productionCostController.GetProductionCostByID)

	productionCostRouter.PUT("/:farmID/lote/:batchID/plantacoes/:plantingID/custos/:costID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("costID"),
		productionCostController.PutProductionCost)

	productionCostRouter.DELETE("/:farmID/lote/:batchID/plantacoes/:plantingID/custos/:costID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("costID"),
		productionCostController.DeleteProductionCost)

}
