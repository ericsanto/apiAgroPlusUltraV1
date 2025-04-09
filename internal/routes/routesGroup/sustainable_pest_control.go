package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterSustainablePestControl(r *gin.Engine) {

	repositorySustainablePestControl := repositories.NewSustainablePestControlRepository(config.DB)
	serviceSustainablePestControl := services.NewSustainablePestControlService(repositorySustainablePestControl)
	controllerSustainablePestControl := controllers.NewSustainablePestControlController(serviceSustainablePestControl)

	routerSustainablePestControl := r.Group("/v1/controle-de-pragas")
	routerSustainablePestControl.GET("/", controllerSustainablePestControl.GetAllSustainablePestControl)
	routerSustainablePestControl.POST("/", controllerSustainablePestControl.PostSustainablePestControl)
	routerSustainablePestControl.GET("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.GetFindByIdSustainablePestControl)
	routerSustainablePestControl.PUT("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.PutSustainablePestControl)
	routerSustainablePestControl.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.DeleteSustainablePestControl)
}
