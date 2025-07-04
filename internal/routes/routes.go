package routes

import (
	"github.com/gin-gonic/gin"

	routesgroup "github.com/ericsanto/apiAgroPlusUltraV1/internal/routes/routesGroup"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()

	routesgroup.SetupRouterAgricultureCulture(router)
	routesgroup.SetupRoutesSoilType(router)
	routesgroup.SetupRoutesTypePest(router)
	routesgroup.SetupRouterPest(router)
	routesgroup.SetupRouterPestAgricultureCulture(router)
	// routesgroup.SetupRouterIrrigationRecomended(router)
	routesgroup.SetupRouterAgricultureCultureIrrigation(router)
	routesgroup.SetupRouterSustainablePestControl(router)
	routesgroup.SetupAgricultureCulturePestMethod(router)
	routesgroup.SetupBatchRouter(router)
	routesgroup.SetupRouterPlanting(router)
	routesgroup.SetupProductionCostRouter(router)
	routesgroup.SetupSalePlantingRouter(router)
	routesgroup.SetupProfitRouter(router)
	routesgroup.SetupPerformancePlantingRouter(router)
	routesgroup.SetupRouterFarm(router)
	routesgroup.SetupRouterDetectPestImage(router)
	routesgroup.RouterGroupDiseaseDetect(router)
	routesgroup.RouterGroupOpenWeather(router)
	routesgroup.SetupRouterIrrigationDeepSeek(router)
	routesgroup.SetupRouterIrrigationType(router)

	return router
}
