package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupBatchRouter(r *gin.Engine) {

	controllerBatch := di.NewBatchBuilder().Builder()

	batchRouter := r.Group("/v1/batchs")
	batchRouter.POST("/", controllerBatch.PostBatch)
	batchRouter.GET("/", controllerBatch.GetAllBatch)
	batchRouter.GET("/:id", middlewares.ValidateIdParam("id"), controllerBatch.GetBatchFindById)
	batchRouter.PUT("/:id", middlewares.ValidateIdParam("id"), controllerBatch.PutBatch)
	batchRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerBatch.DeleteBatch)

}
