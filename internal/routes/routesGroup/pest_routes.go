package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)


func SetupRouterPest(r *gin.Engine) {
  pestRepository := repositories.NewPestRepository(config.DB)
  pestService := services.NewPestService(pestRepository)
  pestController := controllers.NewPestController(pestService)


  pestRouter := r.Group("/v1/pragas")
  pestRouter.GET("/", pestController.GetAllPestController)
  pestRouter.GET("/:id", middlewares.ValidateIdParam("id"), pestController.GetFindByIdPestController)
  pestRouter.POST("/", pestController.PostPestController)
  pestRouter.PUT("/:id", middlewares.ValidateIdParam("id"), pestController.PutPestController)
  pestRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), pestController.DeletePestController)
}

