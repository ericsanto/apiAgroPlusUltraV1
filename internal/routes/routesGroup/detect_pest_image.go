package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouterDetectPestImage(r *gin.Engine) {

	detectPestImageRouter := r.Group("/v1/reconhecimento-de-praga")
	detectPestImageRouter.POST("/", middlewares.ValidateJWT(), controllers.DetectPestImage)
}
