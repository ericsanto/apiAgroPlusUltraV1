package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupPerformancePlantingRouter(r *gin.Engine) {

	repositoryPerformanceCulture := repositories.NewPerformanceCultureRepository(config.DB)
	servicePerformanceCulture := services.NewPerformancePlantingService(repositoryPerformanceCulture)
	controllerPerformanceCulture := controllers.NewPerformancePlantingController(servicePerformanceCulture)

	perfomanceCultureRouter := r.Group("/v1/perfomances-das-plantacoes")

	perfomanceCultureRouter.POST("/", controllerPerformanceCulture.PostPerformanceCulture)
	perfomanceCultureRouter.GET("/", controllerPerformanceCulture.GetAllPerformancePlanting)
	perfomanceCultureRouter.PUT("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.PutPerformancePlanting)
	perfomanceCultureRouter.GET("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.GetPerformancePlantingByID)
	perfomanceCultureRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerPerformanceCulture.DeletePerformancePlanting)
}
