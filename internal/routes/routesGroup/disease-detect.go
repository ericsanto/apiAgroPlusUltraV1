package routesgroup

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RouterGroupDiseaseDetect(r *gin.Engine) {

	routerDiseaseDetectGroup := r.Group("/v1/disease-detect")

	routerDiseaseDetectGroup.POST("/", controllers.DiseaseDetect)
}
