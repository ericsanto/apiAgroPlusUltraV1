package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)


func SetupRoutesSoilType(r *gin.Engine){

  typeSoilRepo := repositories.NewSoilRepository(config.DB)
  typeSoilService := services.NewSoilTypeService(typeSoilRepo)
  typeSoilHandler := controllers.NewSoilTypeController(typeSoilService)

  soilTypeRouterGroup := r.Group("/v1/tipos-de-solo")

  soilTypeRouterGroup.GET("/", typeSoilHandler.GetAllSoilTypes)
  soilTypeRouterGroup.GET("/:id",middlewares.ValidateIdParam("id"), typeSoilHandler.GetSoilTypeFindById)
  soilTypeRouterGroup.POST("/", typeSoilHandler.PostSoilType)
  soilTypeRouterGroup.PUT("/:id", middlewares.ValidateIdParam("id"), typeSoilHandler.PutSoilType)
  soilTypeRouterGroup.DELETE("/:id",middlewares.ValidateIdParam("id"), typeSoilHandler.DeleteSoilType)
 

}
