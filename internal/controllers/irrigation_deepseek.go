package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type IrrigationRecommendedDeepseekController struct {
	irrigationDeepseekService *services.IrrigationRecommendedDeepSeekService
}

func NewIrrigationRecommendedDeepseekController(irrigationDeepseekService *services.IrrigationRecommendedDeepSeekService) *IrrigationRecommendedDeepseekController {
	return &IrrigationRecommendedDeepseekController{irrigationDeepseekService: irrigationDeepseekService}

}

func (i *IrrigationRecommendedDeepseekController) IrrigationDeepseek(c *gin.Context) {

	val, _ := c.Get("lat")

	lat := val.(float64)

	val, _ = c.Get("long")

	long := val.(float64)

	if err := i.irrigationDeepseekService.IrrigationRecommendedDeepSeek(lat, long); err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

}
