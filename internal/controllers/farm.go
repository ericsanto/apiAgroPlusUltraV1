package controllers

import (
	"errors"
	"net/http"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type FamrController struct {
	serviceFarm *services.FarmService
}

func NewFarmController(serviceFarm *services.FarmService) *FamrController {
	return &FamrController{serviceFarm: serviceFarm}
}

func (fc *FamrController) PostFarm(c *gin.Context) {

	var farmRequest requests.FarmRequest

	val, exist := c.Get("userID")
	if !exist {
		return
	}

	userID := val.(uint)

	farmRequest = requests.FarmRequest{
		Name:   farmRequest.Name,
		UserID: userID,
	}

	if err := c.ShouldBindJSON(&farmRequest); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(farmRequest)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := fc.serviceFarm.Create(farmRequest); err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		default:
			myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
			return
		}
	}

	c.Status(http.StatusCreated)

}
