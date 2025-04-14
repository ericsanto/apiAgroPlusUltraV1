package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupAgricultureCulturePestMethod(r *gin.Engine) {

	agricultureCulturePestMethodRepository := repositories.NewAgricultureCulturePestMethodRepository(config.DB)
	agricultureCulturePestMethodService := services.NewAgricultureCulturePestMethodService(agricultureCulturePestMethodRepository)
	agricultureCulturePestMethodController := controllers.NewAgricultureCulturePestMethodController(agricultureCulturePestMethodService)

	routerAgricultureCulturePestMethod := r.Group("/v1/controle-de-praga-agricultura")
	routerAgricultureCulturePestMethod.POST("/", agricultureCulturePestMethodController.PostAgricultureCulturePestMethod)
	routerAgricultureCulturePestMethod.GET("/", middlewares.ValidateQueryParamAgricultureCulturePestMethod(), agricultureCulturePestMethodController.GetAllAgricultureCulturePestMethod)

}
