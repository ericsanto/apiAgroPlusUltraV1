package controllers

import (
	"net/http"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type AgricultureCulturePestMethodController struct {
	serviceAgricultureCulturePestMethod *services.AgricultureCulturePestMethodService
}

func NewAgricultureCulturePestMethodController(serviceAgricultureCulturePestMethod *services.AgricultureCulturePestMethodService) *AgricultureCulturePestMethodController {
	return &AgricultureCulturePestMethodController{serviceAgricultureCulturePestMethod: serviceAgricultureCulturePestMethod}
}

func (a *AgricultureCulturePestMethodController) PostAgricultureCulturePestMethod(c *gin.Context) {

	var requestAgricultureCulturePestMethod requests.AgricultureCulturePestMethodRequest

	if err := c.ShouldBindJSON(&requestAgricultureCulturePestMethod); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestAgricultureCulturePestMethod)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := a.serviceAgricultureCulturePestMethod.PostAgricultureCulturePestMethod(requestAgricultureCulturePestMethod); err != nil {

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)
}

func (a *AgricultureCulturePestMethodController) GetAllAgricultureCulturePestMethod(c *gin.Context) {

	val, _ := c.Get("agricultureCultureName")

	agricultureCultureName := val

	val, _ = c.Get("pestName")

	pestName := val

	val, _ = c.Get("sustainablePestControlMethod")

	sustainablePestControlMethod := val

	if agricultureCultureName != "" || pestName != "" || sustainablePestControlMethod != "" {
		responseAgricultureCulturePestMethod, err := a.serviceAgricultureCulturePestMethod.GetAllAgricultureCulturePestMethodByParam(agricultureCultureName, pestName, sustainablePestControlMethod)
		if err != nil {
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

		c.JSON(http.StatusOK, responseAgricultureCulturePestMethod)
		return
	}

	responseAgricultureCulturePestMethod, err := a.serviceAgricultureCulturePestMethod.GetAllAgricultureCulturePestMethod()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responseAgricultureCulturePestMethod)

}
