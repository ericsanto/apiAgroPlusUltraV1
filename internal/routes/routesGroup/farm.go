package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterFarm(r *gin.Engine) {

	farmController := di.NewFarmBuilder().Builder()

	farmRouterGroup := r.Group("/v1/fazenda")
	farmRouterGroup.POST("/", middlewares.ValidateJWT(), farmController.PostFarm)

}
