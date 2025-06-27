package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterIrrigationDeepSeek(r *gin.Engine) {

	routerIrrigationDeepseekController, _ := di.NewIrrigationDeepSeekBuilder().Builder()
	routerIrrigationDeepseek := r.Group("/v1/irrigation-deepseek")

	routerIrrigationDeepseek.GET("/", middlewares.GetCoordinates(), routerIrrigationDeepseekController.IrrigationDeepseek)

}
