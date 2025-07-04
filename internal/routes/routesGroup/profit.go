package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupProfitRouter(r *gin.Engine) {

	profitController := di.NewProfitBuilder().Builder()

	routerProfit := r.Group("/v1/lucro")
	routerProfit.GET("/:id", middlewares.ValidateIdParam("id"), middlewares.ValidateJWT(), profitController.GetProfit)
}
