package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupSalePlantingRouter(r *gin.Engine) {

	salePlantingController := di.NewSalePlantingBuilder().Builder()

	salePlantingRouter := r.Group("/v1/vendas-plantacoes")
	salePlantingRouter.POST("/", salePlantingController.PostSalePlanting)
	salePlantingRouter.GET("/", salePlantingController.GetAllSalePlanting)
	salePlantingRouter.GET("/:id", middlewares.ValidateIdParam("id"), salePlantingController.GetSalePlantingByID)
	salePlantingRouter.PUT("/:id", middlewares.ValidateIdParam("id"), salePlantingController.PutSalePlanting)
	salePlantingRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), salePlantingController.DeleteSalePlanting)
}
