package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterPestAgricultureCulture(r *gin.Engine) {

	pestAgricultureCultureRepository := repositories.NewPestAgricultureCultureRepository(config.DB)
	pestAgricultureCultureService := services.NewPestAgricultureCultureService(pestAgricultureCultureRepository)
	pestAgricultureCultureController := controllers.NewPestAgricultureCultureController(pestAgricultureCultureService)

	pestAgricultureCultureRouter := r.Group("/v1/pragas-das-culturas-agricolas")
	pestAgricultureCultureRouter.GET("/", pestAgricultureCultureController.GetAllAgricultureCultureController)
	pestAgricultureCultureRouter.GET("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.GetFindByIdAgricultureCultureController)
	pestAgricultureCultureRouter.POST("/", pestAgricultureCultureController.PostPestAgricultureCultureController)
	pestAgricultureCultureRouter.PUT("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.PutPestAgricultureCulture)
	pestAgricultureCultureRouter.DELETE("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.DeletePestAgricultureCultureController)
}
