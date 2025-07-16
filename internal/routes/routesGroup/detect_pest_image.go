package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRouterDetectPestImage(r *gin.Engine) {

	builderController, _ := di.NewPestDetectImageBuilder().Builder()

	detectPestImageRouter := r.Group("/v1/pragas/reconhecimentos")
	detectPestImageRouter.POST("/", middlewares.ValidateJWT(), builderController.DetectPestImage)
}
