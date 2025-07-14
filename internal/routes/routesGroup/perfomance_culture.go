package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupPerformancePlantingRouter(r *gin.Engine) {

	controllerPerformanceCulture := di.NewPerfomancePlantingBuilder().Builder()

	perfomanceCultureRouter := r.Group("/v1/fazenda")

	perfomanceCultureRouter.POST("/:farmID/lote/:batchID/plantacoes/:plantingID/performances",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		controllerPerformanceCulture.PostPerformanceCulture)

	perfomanceCultureRouter.GET("/:farmID/lote/:batchID/plantacoes/performances",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		controllerPerformanceCulture.GetAllPerformancePlanting)

	perfomanceCultureRouter.PUT("/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("performanceID"),
		controllerPerformanceCulture.PutPerformancePlanting)

	perfomanceCultureRouter.GET("/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("performanceID"),
		controllerPerformanceCulture.GetPerformancePlantingByID)

	perfomanceCultureRouter.DELETE("/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		middlewares.ValidateIdParam("performanceID"),
		controllerPerformanceCulture.DeletePerformancePlanting)

}
