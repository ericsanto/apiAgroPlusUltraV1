package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RouterGroupDiseaseDetect(r *gin.Engine) {

	routerDiseaseDetectGroup := r.Group("/v1/disease-detect")

	routerDiseaseDetectGroup.POST("/", middlewares.ValidateJWT(), controllers.DiseaseDetect)
}
