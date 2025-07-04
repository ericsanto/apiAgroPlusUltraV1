package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func RouterGroupDiseaseDetect(r *gin.Engine) {

	diseaseDetectContorller, _ := di.NewDiseaseDetectBuilder().Builder()

	routerDiseaseDetectGroup := r.Group("/v1/disease-detect")

	routerDiseaseDetectGroup.POST("/", middlewares.ValidateJWT(), diseaseDetectContorller.DiseaseDetect)
}
