package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type ProfitControllerInterface interface {
	GetProfit(c *gin.Context)
}

type ProfitController struct {
	profitService services.ProfitServiceInterface
}

func NewProfitController(profitService services.ProfitServiceInterface) ProfitControllerInterface {
	return &ProfitController{profitService: profitService}
}

func (p *ProfitController) GetProfit(c *gin.Context) {

	val, exists := c.Get("userID")

	if !exists {
		return
	}

	userID := val.(uint)

	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	responseProfit, err := p.profitService.GetProfit(batchID, farmID, userID, plantingID)
	if err != nil {

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			myerror.HttpErrors(http.StatusNotFound, "não existe  plantação com id", c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.JSON(http.StatusOK, responseProfit)
}
