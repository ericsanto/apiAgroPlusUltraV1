package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type IrrigationTypeController struct {
	irrigationTypeService *services.IrrigationTypeService
}

func NewIrrigationTypeController(irrigationTypeService *services.IrrigationTypeService) *IrrigationTypeController {
	return &IrrigationTypeController{irrigationTypeService: irrigationTypeService}
}

func (it *IrrigationTypeController) GetAllIrrigationType(c *gin.Context) {

	IrrigationType, err := it.irrigationTypeService.GetAllIrrigationType()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, IrrigationType)
}

func (it *IrrigationTypeController) PostIrrigationType(c *gin.Context) {

	var requesIrrigationType requests.IrrigationTypeRequest

	if err := c.ShouldBindJSON(&requesIrrigationType); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requesIrrigationType)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(valid) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, valid, c)
		return
	}

	if err := it.irrigationTypeService.PostirrigationType(requesIrrigationType); err != nil {
		log.Println(err.Error())

		if errors.Is(err, myerror.ErrEnumInvalid) {
			myerror.HttpErrors(http.StatusUnprocessableEntity, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro ao criar tipo de irrigacao", c)
		return
	}

	c.Status(http.StatusCreated)
}

func (it *IrrigationTypeController) GetIrrigationTypeByID(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	irrigationType, err := it.irrigationTypeService.GetIrrigationTypeByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "não existe tipo de irrigacao com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe tipo de irrigacao com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, irrigationType)
}

func (it *IrrigationTypeController) PutIrrigationType(c *gin.Context) {

	var requesIrrigationType requests.IrrigationTypeRequest

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := c.ShouldBindJSON(&requesIrrigationType); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requesIrrigationType)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(valid) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, valid, c)
		return
	}

	if err := it.irrigationTypeService.PutIrrigationType(id, requesIrrigationType); err != nil {
		log.Println(err.Error())
		myerror.HttpErrors(http.StatusInternalServerError, "erro ao atualizar tipo de irragacao", c)
		return
	}

	c.Status(http.StatusOK)

}

func (it *IrrigationTypeController) DeleteIrrigationType(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := it.irrigationTypeService.DeleteIrrigationType(id); err != nil {
		if strings.Contains(err.Error(), "não existe tipo de irrigacao com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe tipo de irrigacao com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
