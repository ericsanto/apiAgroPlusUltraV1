package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type PerfomancePlantingController struct {
	servvicePerfomancePlanting *services.PerfomancePlantingService
}

func NewPerfomancePlantingController(servvicePerfomancePlanting *services.PerfomancePlantingService) *PerfomancePlantingController {
	return &PerfomancePlantingController{servvicePerfomancePlanting: servvicePerfomancePlanting}
}

func (p *PerfomancePlantingController) PostPerfomanceCulture(c *gin.Context) {

	var requestPerfomanceCulture requests.PerfomancePlantingRequest

	if err := c.ShouldBindJSON(&requestPerfomanceCulture); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPerfomanceCulture)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := p.servvicePerfomancePlanting.PostPerfomancePlanting(requestPerfomanceCulture); err != nil {
		switch {
		case errors.Is(err, myerror.ErrDuplicateKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrViolatedForeingKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrEnumInvalid):
			myerror.HttpErrors(http.StatusUnprocessableEntity, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			fmt.Println(err.Error())
			return
		}

	}

	c.Status(http.StatusCreated)
}

func (p *PerfomancePlantingController) GetAllPerfomancePlanting(c *gin.Context) {

	perfomancePlanting, err := p.servvicePerfomancePlanting.GetAllPerfomancePlanting()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, perfomancePlanting)
}

func (p *PerfomancePlantingController) PutPerfomancePlanting(c *gin.Context) {

	var requestPerfomancePlanting requests.PerfomancePlantingRequest

	if err := c.ShouldBindJSON(&requestPerfomancePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPerfomancePlanting)

	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := p.servvicePerfomancePlanting.PutPerformancePlanting(id, requestPerfomancePlanting); err != nil {
		switch {
		case errors.Is(err, myerror.ErrDuplicateKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrViolatedForeingKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.Status(http.StatusOK)

}

func (p *PerfomancePlantingController) GetPerformancePlantingByID(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	responsePerfomancePlanting, err := p.servvicePerfomancePlanting.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(id)
	if err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.JSON(http.StatusOK, responsePerfomancePlanting)
}

func (p *PerfomancePlantingController) DeletePerfomancePlanting(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := p.servvicePerfomancePlanting.DeletePerfomancePlanting(id); err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.Status(http.StatusNoContent)
}
