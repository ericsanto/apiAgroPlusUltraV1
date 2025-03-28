package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)


func SetupRoutesTypePest(r *gin.Engine) {

  typePestRepository := repositories.NewTypePestRepository(config.DB)
  typePestService := services.NewTypePestService(typePestRepository)
  typePestController := controllers.NewTypePestController(typePestService) 

  typePests := r.Group("/v1/pragas")
  typePests.GET("/", typePestController.GetAllTypePestController)
  typePests.GET("/:id", middlewares.ValidateIdParam("id"), typePestController.GetAllTypePestFindByIdController)
  typePests.POST("/", typePestController.PostTypePestController)
  typePests.PUT("/:id", middlewares.ValidateIdParam("id"), typePestController.PutTypePestController)
  typePests.DELETE("/:id", middlewares.ValidateIdParam("id"), typePestController.DeleteTypePestController)
}
