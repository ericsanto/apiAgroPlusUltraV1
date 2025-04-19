package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupSalePlantingRouter(r *gin.Engine) {

	salePlantingRepository := repositories.NewSalePlantingRepository(config.DB)
	salePlantingService := services.NewSalePlantingService(salePlantingRepository)
	salePlantingController := controllers.NewSalePlantingController(salePlantingService)

	salePlantingRouter := r.Group("/v1/vendas-plantacoes")
	salePlantingRouter.POST("/", salePlantingController.PostSalePlanting)
	salePlantingRouter.GET("/", salePlantingController.GetAllSalePlanting)
	salePlantingRouter.GET("/:id", middlewares.ValidateIdParam("id"), salePlantingController.GetSalePlantingByID)
	salePlantingRouter.PUT("/:id", middlewares.ValidateIdParam("id"), salePlantingController.PutSalePlanting)
	salePlantingRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), salePlantingController.DeleteSalePlanting)
}
