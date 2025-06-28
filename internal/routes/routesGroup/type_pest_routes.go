package routesgroup

import (
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/di"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/middlewares"
)

func SetupRoutesTypePest(r *gin.Engine) {

	typePestsController := di.NewTypePestBuilder().Builder()

	typePests := r.Group("/v1/tipos-de-pragas")
	typePests.GET("/", typePestsController.GetAllTypePestController)
	typePests.GET("/:id", middlewares.ValidateIdParam("id"), typePestsController.GetAllTypePestFindByIdController)
	typePests.POST("/", typePestsController.PostTypePestController)
	typePests.PUT("/:id", middlewares.ValidateIdParam("id"), typePestsController.PutTypePestController)
	typePests.DELETE("/:id", middlewares.ValidateIdParam("id"), typePestsController.DeleteTypePestController)
}
