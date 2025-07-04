package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type DiseaseDetectControllerInterface interface {
	DiseaseDetect(c *gin.Context)
}

type DiseaseDetectController struct {
	diseaseDetectService services.DiseaseDetectServiceInterface
}

func NewDiseaseDetectController(diseaseDetectService services.DiseaseDetectServiceInterface) DiseaseDetectControllerInterface {
	return &DiseaseDetectController{diseaseDetectService: diseaseDetectService}
}

func (ddc *DiseaseDetectController) DiseaseDetect(c *gin.Context) {

	ctx := context.Background()

	responseApiPython, err := ddc.diseaseDetectService.ServiceDiseaseDetect(c.Request)
	fmt.Println(responseApiPython)

	if err != nil {
		switch {
		case errors.Is(err, ctx.Err()):
			myerror.HttpErrors(http.StatusRequestTimeout, err.Error(), c)
			return
		case errors.As(err, &sarama.ConsumerError{}):
			log.Println(err.Error())
			myerror.HttpErrors(http.StatusBadGateway, "erro ao consultar servidor externo", c)
			return
		case errors.Is(err, myerror.ErrImageSizeToLarge):
			myerror.HttpErrors(413, err.Error(), c)
			return
		case errors.Is(err, myerror.ErrUnsupportedImageType):
			myerror.HttpErrors(http.StatusUnsupportedMediaType, err.Error(), c)
			return
		default:
			myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
			return
		}
	}

	c.JSON(http.StatusOK, responseApiPython)
}
