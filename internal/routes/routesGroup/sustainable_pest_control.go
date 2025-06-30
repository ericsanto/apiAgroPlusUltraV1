package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterSustainablePestControl(r *gin.Engine) {

	controllerSustainablePestControl := di.NewSustainablePestControl().Builder()

	routerSustainablePestControl := r.Group("/v1/controle-de-pragas")
	routerSustainablePestControl.GET("/", controllerSustainablePestControl.GetAllSustainablePestControl)
	routerSustainablePestControl.POST("/", controllerSustainablePestControl.PostSustainablePestControl)
	routerSustainablePestControl.GET("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.GetFindByIdSustainablePestControl)
	routerSustainablePestControl.PUT("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.PutSustainablePestControl)
	routerSustainablePestControl.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerSustainablePestControl.DeleteSustainablePestControl)
}
