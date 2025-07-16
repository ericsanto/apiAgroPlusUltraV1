package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterPlanting(r *gin.Engine) {

	plantingController := di.NewPlantingRepository().Builder()

	plantingRouter := r.Group("/v1/fazenda")
	plantingRouter.POST("/:farmID/lote/:batchID/plantacoes",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		plantingController.PostPlanting)

	plantingRouter.PUT("/:farmID/lote/:batchID/plantacoes/:plantingID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		plantingController.PutPlanting)

	plantingRouter.GET("/:farmID/lote/:batchID/plantacoes",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateQueryParamPlanting(),
		plantingController.GetPlantingQueryParamBatchNameOrActive)

	plantingRouter.DELETE("/:farmID/lote/:batchID/plantacoes/:plantingID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		plantingController.DeletePlanting)
}
