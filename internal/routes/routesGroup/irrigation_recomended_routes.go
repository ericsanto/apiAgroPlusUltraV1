package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouterIrrigationRecomended(r *gin.Engine) {
	repositoryIrrigationRecomended := repositories.NewIrrigationRecomdedRepository(config.DB)
	serviceIrrigationRecomended := services.NewIrrigationRecomendedService(repositoryIrrigationRecomended)
	controllerIrrigationRecomended := controllers.NewIrrigationRecomendedController(serviceIrrigationRecomended)

	irrigationRouter := r.Group("/v1/irrigacao")
	irrigationRouter.GET("/", controllerIrrigationRecomended.GetAllIrrigationRecomended)
	irrigationRouter.POST("/", controllerIrrigationRecomended.PostIrrigationRecomended)
	irrigationRouter.GET("/:id", middlewares.ValidateIdParam("id"), controllerIrrigationRecomended.GetByIdrrigationRecomended)
	irrigationRouter.PUT("/:id", middlewares.ValidateIdParam("id"), controllerIrrigationRecomended.PutIrrigationRecomendedController)
	irrigationRouter.DELETE("/:id", middlewares.ValidateIdParam("id"), controllerIrrigationRecomended.DeleteIrrigationRecomendedController)
}
