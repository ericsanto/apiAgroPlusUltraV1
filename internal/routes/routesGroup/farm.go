package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterFarm(r *gin.Engine) {

	farmRepository := repositories.NewFarmRepository(config.DB)
	farmService := services.NewFarmService(farmRepository)
	farmController := controllers.NewFarmController(farmService)

	farmRouterGroup := r.Group("/v1/fazenda")
	farmRouterGroup.POST("/", middlewares.ValidateJWT(), farmController.PostFarm)

}
