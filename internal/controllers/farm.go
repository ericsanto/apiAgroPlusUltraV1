package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type FarmControllerInterface interface {
	PostFarm(c *gin.Context)
	GetFarmByID(c *gin.Context)
	GetAllFarm(c *gin.Context)
}

type FarmController struct {
	serviceFarm services.FarmServiceInterface
}

func NewFarmController(serviceFarm services.FarmServiceInterface) FarmControllerInterface {
	return &FarmController{serviceFarm: serviceFarm}
}

func (fc *FarmController) PostFarm(c *gin.Context) {

	ctx := c.Request.Context()

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

	if err := fc.serviceFarm.Create(ctx, farmRequest); err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusConflict, fmt.Sprintf("%s ja existe fazenda com nome fornecido", err.Error()), c)
			return
		}
	}

	c.Status(http.StatusCreated)

}

func (fc *FarmController) GetFarmByID(c *gin.Context) {

	val, exist := c.Get("userID")
	if !exist {
		return
	}

	userID := val.(uint)

	validateID := validators.GetAndValidateIdMidlware(c, "id")

	responseFarm, err := fc.serviceFarm.GetFarmByID(userID, validateID)

	if err != nil {
		switch {
		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.JSON(http.StatusOK, responseFarm)
}

func (fc *FarmController) GetAllFarm(c *gin.Context) {

	val, exists := c.Get("userID")

	if !exists {
		return
	}

	userID := val.(uint)

	listFarmResponse, err := fc.serviceFarm.GetAllFarm(userID)

	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, listFarmResponse)
}
