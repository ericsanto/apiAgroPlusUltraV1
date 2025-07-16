package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupProfitRouter(r *gin.Engine) {

	profitController := di.NewProfitBuilder().Builder()

	routerProfit := r.Group("/v1/fazenda")
	routerProfit.GET("/:farmID/lote/:batchID/plantacoes/:plantingID/lucros",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		profitController.GetProfit)
}
