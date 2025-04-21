package controllers

import (
	"errors"
	"net/http"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProfitController struct {
	profitService *services.ProfitService
}

func NewProfitController(profitService *services.ProfitService) *ProfitController {
	return &ProfitController{profitService: profitService}
}

func (p *ProfitController) GetProfit(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	responseProfit, err := p.profitService.GetProfit(id)
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
