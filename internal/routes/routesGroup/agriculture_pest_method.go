package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupAgricultureCulturePestMethod(r *gin.Engine) {

	agricultureCulturePestMethodController := di.NewAgricultureCulturePestMethodBuilder().Builder()

	routerAgricultureCulturePestMethod := r.Group("/v1/controle-de-praga-agricultura")
	routerAgricultureCulturePestMethod.POST("/", agricultureCulturePestMethodController.PostAgricultureCulturePestMethod)
	routerAgricultureCulturePestMethod.GET("/", middlewares.ValidateQueryParamAgricultureCulturePestMethod(), agricultureCulturePestMethodController.GetAllAgricultureCulturePestMethod)

}
