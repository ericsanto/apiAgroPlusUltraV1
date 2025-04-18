package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type PlantingController struct {
	plantingService *services.PlantingService
}

func NewPlantingController(plantingService *services.PlantingService) *PlantingController {
	return &PlantingController{plantingService: plantingService}
}

func (p *PlantingController) PostPlanting(c *gin.Context) {

	var requestPlanting requests.PlantingRequest

	if err := c.ShouldBindJSON(&requestPlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := p.plantingService.PostPlanting(requestPlanting); err != nil {
		if strings.Contains(err.Error(), "erro ao cadastrar plantação. Lote já está sendo utilizado pela cultura") {
			myerror.HttpErrors(http.StatusConflict, "erro ao cadastrar objeto. Lote já esta sendo utilizado", c)
			return
		}
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)
}

func (p *PlantingController) GetPlantingQueryParamBatchNameOrActive(c *gin.Context) {

	val, exist := c.Get("batchName")
	if !exist {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	batchName := val

	val, exist = c.Get("active")
	if !exist {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	active := val.(string)

	if batchName == "" && active == "" {
		responsePlanting, err := p.plantingService.GetAllPlanting()
		if err != nil {
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

		c.JSON(http.StatusOK, responsePlanting)
		return
	}

	activeBool, err := strconv.ParseBool(active)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	responsePlanting, err := p.plantingService.GetByParamBatchNameOrIsActivePlanting(batchName.(string), activeBool)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responsePlanting)
}

func (p *PlantingController) PutPlanting(c *gin.Context) {

	var requestPlanting requests.PlantingRequest

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := c.ShouldBindJSON(&requestPlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := p.plantingService.PutPlanting(id, requestPlanting); err != nil {
		if strings.Contains(err.Error(), "não existe plantação com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe plantação com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)

}

func (p *PlantingController) DeletePlanting(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := p.plantingService.DeletePlanting(id); err != nil {
		if strings.Contains(err.Error(), "não existe plantação com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe plantação com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
