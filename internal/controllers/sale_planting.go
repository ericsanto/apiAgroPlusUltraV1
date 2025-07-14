package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type SalePlantingControllerInterface interface {
	PostSalePlanting(c *gin.Context)
	GetAllSalePlanting(c *gin.Context)
	GetSalePlantingByID(c *gin.Context)
	PutSalePlanting(c *gin.Context)
	DeleteSalePlanting(c *gin.Context)
}

type SalePlantingController struct {
	salePlantingService services.SalePlantingServiceInterface
}

func NewSalePlantingController(salePlantingService services.SalePlantingServiceInterface) SalePlantingControllerInterface {
	return &SalePlantingController{salePlantingService: salePlantingService}
}

func (s *SalePlantingController) PostSalePlanting(c *gin.Context) {

	var requestSalePlanting requests.SalePlantingRequest

	if err := c.ShouldBindJSON(&requestSalePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSalePlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	if err := s.salePlantingService.PostSalePlanting(batchID, farmID, userID, plantingID, requestSalePlanting); err != nil {

		switch {
		case errors.Is(err, myerror.ErrDuplicateSale):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)

}

func (s *SalePlantingController) GetAllSalePlanting(c *gin.Context) {

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	responsesSalePlanting, err := s.salePlantingService.GetAllSalePlanting(batchID, farmID, userID)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responsesSalePlanting)
}

func (s *SalePlantingController) GetSalePlantingByID(c *gin.Context) {

	salePlantingID := validators.GetAndValidateIdMidlware(c, "salePlantingID")
	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	responseSalePlanting, err := s.salePlantingService.GetSalePlantingByID(batchID, farmID, userID, salePlantingID)
	if err != nil {
		if errors.Is(err, myerror.ErrNotFoundSale) {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		if errors.Is(err, myerror.ErrNotFound) {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		if errors.Is(err, myerror.ErrFarmNotFound) {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responseSalePlanting)
}

func (s *SalePlantingController) PutSalePlanting(c *gin.Context) {

	var requestSalePlanting requests.SalePlantingRequest

	if err := c.ShouldBindJSON(&requestSalePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSalePlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	salePlantingID := validators.GetAndValidateIdMidlware(c, "salePlantingID")
	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if err := s.salePlantingService.PutSalePlanting(batchID, farmID, userID, salePlantingID, requestSalePlanting); err != nil {

		switch {
		case errors.Is(err, myerror.ErrDuplicateSale):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFoundSale):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrViolatedForeingKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.Status(http.StatusOK)
}

func (s *SalePlantingController) DeleteSalePlanting(c *gin.Context) {

	salePlantingID := validators.GetAndValidateIdMidlware(c, "salePlantingID")
	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if err := s.salePlantingService.DeleteSalePlanting(batchID, farmID, userID, salePlantingID); err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFoundSale):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.Status(http.StatusNoContent)
}
