package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupSalePlantingRouter(r *gin.Engine) {

	salePlantingController := di.NewSalePlantingBuilder().Builder()

	salePlantingRouter := r.Group("/v1/fazenda")
	salePlantingRouter.POST("/:farmID/lote/:batchID/plantacoes/:plantingID/vendas",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("plantingID"),
		salePlantingController.PostSalePlanting)

	salePlantingRouter.GET("/:farmID/lote/:batchID/plantacoes/vendas",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		salePlantingController.GetAllSalePlanting)

	salePlantingRouter.GET("/:farmID/lote/:batchID/plantacoes/vendas/:salePlantingID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("salePlantingID"),
		salePlantingController.GetSalePlantingByID)

	salePlantingRouter.PUT("/:farmID/lote/:batchID/plantacoes/vendas/:salePlantingID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("salePlantingID"),
		salePlantingController.PutSalePlanting)

	salePlantingRouter.DELETE("/:farmID/lote/:batchID/plantacoes/vendas/:salePlantingID",
		middlewares.ValidateJWT(),
		middlewares.ValidateIdParam("farmID"),
		middlewares.ValidateIdParam("batchID"),
		middlewares.ValidateIdParam("salePlantingID"),
		salePlantingController.DeleteSalePlanting)
}
