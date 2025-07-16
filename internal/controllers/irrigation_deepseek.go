package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type IrrigationRecommendedDeepSeekControllerInterface interface {
	IrrigationDeepseek(c *gin.Context)
}

type IrrigationRecommendedDeepseekController struct {
	irrigationDeepseekService services.IrrigationRecommendedDeepSeekServiceInterface
}

func NewIrrigationRecommendedDeepseekController(irrigationDeepseekService services.IrrigationRecommendedDeepSeekServiceInterface) IrrigationRecommendedDeepSeekControllerInterface {
	return &IrrigationRecommendedDeepseekController{irrigationDeepseekService: irrigationDeepseekService}

}

func (i *IrrigationRecommendedDeepseekController) IrrigationDeepseek(c *gin.Context) {

	ctx := c.Request.Context()

	val, _ := c.Get("lat")

	lat := val.(float64)

	val, _ = c.Get("long")

	long := val.(float64)

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")

	if err := i.irrigationDeepseekService.IrrigationRecommendedDeepSeek(ctx, lat, long, userID, farmID); err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.Status(http.StatusOK)

}
