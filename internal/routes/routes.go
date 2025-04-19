package routes

import (
	routesgroup "github.com/ericsanto/apiAgroPlusUltraV1/internal/routes/routesGroup"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()

	routesgroup.SetupRouterAgricultureCulture(router)
	routesgroup.SetupRoutesSoilType(router)
	routesgroup.SetupRoutesTypePest(router)
	routesgroup.SetupRouterPest(router)
	routesgroup.SetupRouterPestAgricultureCulture(router)
	routesgroup.SetupRouterIrrigationRecomended(router)
	routesgroup.SetupRouterAgricultureCultureIrrigation(router)
	routesgroup.SetupRouterSustainablePestControl(router)
	routesgroup.SetupAgricultureCulturePestMethod(router)
	routesgroup.SetupBatchRouter(router)
	routesgroup.SetupRouterPlanting(router)
	routesgroup.SetupProductionCostRouter(router)
	routesgroup.SetupSalePlantingRouter(router)

	return router
}
