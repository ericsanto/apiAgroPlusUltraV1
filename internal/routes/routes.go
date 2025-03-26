package routes

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)


func Routes() *gin.Engine {

  typeSoilRepo := repositories.NewSoilRepository(config.DB)
  typeSoilService := services.NewSoilTypeService(typeSoilRepo)
  typeSoilHandler := controllers.NewSoilTypeController(typeSoilService)
  agricultureCultureRepository := repositories.NewAgricultureCultureRepository(config.DB)
  agricultureCultureService := services.NewAgricultureCultureService(agricultureCultureRepository)
  agricultureCultureHandler := controllers.NewAgricultureController(agricultureCultureService)

  router := gin.Default()

  v1 := router.Group("/v1")
  v1.GET("/tipos-de-solo", typeSoilHandler.GetAllSoilTypes)
  v1.GET("/tipos-de-solo/:id", typeSoilHandler.GetSoilTypeFindById)
  v1.POST("/tipos-de-solo", typeSoilHandler.PostSoilType)
  v1.PUT("/tipos-de-solo/:id", middlewares.ValidateIdParam("id"), typeSoilHandler.PutSoilType)
  v1.DELETE("/tipos-de-solo/:id",middlewares.ValidateIdParam("id"), typeSoilHandler.DeleteSoilType)
  v1.GET("/culturas-agricolas", agricultureCultureHandler.GetAllAgriculturesCultures)
  v1.POST("/culturas-agricolas", agricultureCultureHandler.PostAgricultureCulture)
  return router
}
