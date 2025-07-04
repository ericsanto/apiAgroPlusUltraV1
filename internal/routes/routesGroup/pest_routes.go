package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterPest(r *gin.Engine) {

	pestController := di.NewPestBuilder().Builder()

	pestRouter := r.Group("/v1/pragas")
	pestRouter.GET("/", pestController.GetAllPestController)
	pestRouter.GET("/:id", middlewares.ValidateIdParam("id"), pestController.GetFindByIdPestController)
	pestRouter.POST("/", pestController.PostPestController)
	pestRouter.PUT("/:id", middlewares.ValidateIdParam("id"), pestController.PutPestController)
	pestRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), pestController.DeletePestController)
}
