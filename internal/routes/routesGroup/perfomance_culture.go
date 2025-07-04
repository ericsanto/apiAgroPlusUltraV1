package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupPerformancePlantingRouter(r *gin.Engine) {

	controllerPerformanceCulture := di.NewPerfomancePlantingBuilder().Builder()

	perfomanceCultureRouter := r.Group("/v1/perfomances-das-plantacoes")

	perfomanceCultureRouter.POST("/", controllerPerformanceCulture.PostPerformanceCulture)
	perfomanceCultureRouter.GET("/", controllerPerformanceCulture.GetAllPerformancePlanting)
	perfomanceCultureRouter.PUT("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.PutPerformancePlanting)
	perfomanceCultureRouter.GET("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.GetPerformancePlantingByID)
	perfomanceCultureRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.DeletePerformancePlanting)
}
