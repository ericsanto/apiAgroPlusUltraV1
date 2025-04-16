package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupBatchRouter(r *gin.Engine) {
	respositoryBatch := repositories.NewBatchRepository(config.DB)
	serviceBatch := services.NewBatchService(respositoryBatch)
	controllerBatch := controllers.NewBatchController(serviceBatch)

	batchRouter := r.Group("/v1/batchs")
	batchRouter.POST("/", controllerBatch.PostBatch)
	batchRouter.GET("/", controllerBatch.GetAllBatch)
	batchRouter.GET("/:id", middlewares.ValidateIdParam("id"), controllerBatch.GetBatchFindById)
	batchRouter.PUT("/:id", middlewares.ValidateIdParam("id"), controllerBatch.PutBatch)
	batchRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerBatch.DeleteBatch)

}
