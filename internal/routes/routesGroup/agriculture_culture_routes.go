package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)


func SetupRouterAgricultureCulture(r *gin.Engine){

  agricultureCultureRepository := repositories.NewAgricultureCultureRepository(config.DB)
  agricultureCultureService := services.NewAgricultureCultureService(agricultureCultureRepository)
  agricultureCultureHandler := controllers.NewAgricultureController(agricultureCultureService)

  setupRouterAgricultureCulture := r.Group("/v1/culturas-agricolas")

  setupRouterAgricultureCulture.GET("/", agricultureCultureHandler.GetAllAgriculturesCulturesController)
  setupRouterAgricultureCulture.POST("/", agricultureCultureHandler.PostAgricultureCultureController)
  setupRouterAgricultureCulture.PUT("/:id", middlewares.ValidateIdParam("id"), agricultureCultureHandler.PutAgricultureCultureController)
  setupRouterAgricultureCulture.DELETE("/:id", middlewares.ValidateIdParam("id"), agricultureCultureHandler.DeleteAgricultureCultureController)
 
}
