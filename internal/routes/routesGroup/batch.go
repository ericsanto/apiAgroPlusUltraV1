package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupBatchRouter(r *gin.Engine) {

	controllerBatch := di.NewBatchBuilder().Builder()

	farmRouterGroup := r.Group("/v1/fazenda")

	farmRouterGroup.POST("/:farmID/lote", middlewares.ValidateJWT(), middlewares.ValidateIdParam("farmID"), controllerBatch.PostBatch)
	farmRouterGroup.GET("/:farmID/lote", middlewares.ValidateJWT(), middlewares.ValidateIdParam("farmID"), controllerBatch.GetAllBatch)
	farmRouterGroup.GET("/:farmID/lote/:batchID", middlewares.ValidateJWT(), middlewares.ValidateIdParam("farmID"), middlewares.ValidateIdParam("batchID"), controllerBatch.GetBatchFindById)
	farmRouterGroup.PUT("/:farmID/lote/:batchID", middlewares.ValidateJWT(), middlewares.ValidateIdParam("farmID"), middlewares.ValidateIdParam("batchID"), controllerBatch.PutBatch)
	farmRouterGroup.DELETE("/:farmID/lote/:batchID", middlewares.ValidateJWT(), middlewares.ValidateIdParam("farmID"), middlewares.ValidateIdParam("batchID"), controllerBatch.DeleteBatch)

}
