package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterAgricultureCulture(r *gin.Engine) {

	agricultureCultureHandler := di.NewAgricultureCultureBuiler().Builder()

	setupRouterAgricultureCulture := r.Group("/v1/culturas-agricolas")

	setupRouterAgricultureCulture.GET("/", agricultureCultureHandler.GetAllAgriculturesCulturesController)
	setupRouterAgricultureCulture.POST("/", agricultureCultureHandler.PostAgricultureCultureController)
	setupRouterAgricultureCulture.PUT("/:id", middlewares.ValidateIdParam("id"), agricultureCultureHandler.PutAgricultureCultureController)
	setupRouterAgricultureCulture.DELETE("/:id", middlewares.ValidateIdParam("id"), agricultureCultureHandler.DeleteAgricultureCultureController)

}
