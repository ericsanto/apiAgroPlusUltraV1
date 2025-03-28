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

  return router
}
