package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRoutesSoilType(r *gin.Engine) {

	soilTypeController := di.NewSoilTypeBuilder().Builder()

	soilTypeRouterGroup := r.Group("/v1/tipos-de-solo")

	soilTypeRouterGroup.GET("/", soilTypeController.GetAllSoilTypes)
	soilTypeRouterGroup.GET("/:id", middlewares.ValidateIdParam("id"), soilTypeController.GetSoilTypeFindById)
	soilTypeRouterGroup.POST("/", soilTypeController.PostSoilType)
	soilTypeRouterGroup.PUT("/:id", middlewares.ValidateIdParam("id"), soilTypeController.PutSoilType)
	soilTypeRouterGroup.DELETE("/:id", middlewares.ValidateIdParam("id"), soilTypeController.DeleteSoilType)

}
