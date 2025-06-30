package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterPestAgricultureCulture(r *gin.Engine) {

	pestAgricultureCultureController := di.NewPestAgricultureCultureBuilder().Builder()

	pestAgricultureCultureRouter := r.Group("/v1/pragas-das-culturas-agricolas")
	pestAgricultureCultureRouter.GET("/", pestAgricultureCultureController.GetAllAgricultureCultureController)
	pestAgricultureCultureRouter.GET("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.GetFindByIdAgricultureCultureController)
	pestAgricultureCultureRouter.POST("/", pestAgricultureCultureController.PostPestAgricultureCultureController)
	pestAgricultureCultureRouter.PUT("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.PutPestAgricultureCulture)
	pestAgricultureCultureRouter.DELETE("/relacao", middlewares.ValidateQueryParamPestAgricultureCulture(), pestAgricultureCultureController.DeletePestAgricultureCultureController)
}
