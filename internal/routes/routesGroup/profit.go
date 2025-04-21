package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupProfitRouter(r *gin.Engine) {

	repository := repositories.NewProfitRepository(config.DB)
	service := services.NewProfitService(repository)
	controller := controllers.NewProfitController(service)

	routerProfit := r.Group("/v1/lucro")
	routerProfit.GET("/:id", middlewares.ValidateIdParam("id"), controller.GetProfit)
}
