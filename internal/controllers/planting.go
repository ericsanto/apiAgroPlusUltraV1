package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type PlantingControllerInterface interface {
	PostPlanting(c *gin.Context)
	GetPlantingQueryParamBatchNameOrActive(c *gin.Context)
	PutPlanting(c *gin.Context)
	DeletePlanting(c *gin.Context)
}

type PlantingController struct {
	plantingService services.PlantingServiceInterface
}

func NewPlantingController(plantingService services.PlantingServiceInterface) PlantingControllerInterface {
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if err := p.plantingService.PostPlanting(userID, farmID, batchID, requestPlanting); err != nil {
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if batchName == "" && active == "" {
		responsePlanting, err := p.plantingService.GetAllPlanting(batchID, farmID, userID)
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

	responsePlanting, err := p.plantingService.GetByParamBatchNameOrIsActivePlanting(batchName.(string), activeBool, userID, farmID)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responsePlanting)
}

func (p *PlantingController) PutPlanting(c *gin.Context) {

	var requestPlanting requests.PlantingRequest

	id := validators.GetAndValidateIdMidlware(c, "id")

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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	if err := p.plantingService.PutPlanting(batchID, farmID, userID, plantingID, requestPlanting); err != nil {
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	if err := p.plantingService.DeletePlanting(batchID, farmID, userID, plantingID); err != nil {
		if strings.Contains(err.Error(), "não existe plantação com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe plantação com o id %d", plantingID), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
