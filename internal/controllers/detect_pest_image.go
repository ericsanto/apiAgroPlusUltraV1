package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type DetectPestImageControllerInterface interface {
	DetectPestImage(c *gin.Context)
}

type DetectPestImageController struct {
	DetectPestImageService services.DetectPestImageServiceInterface
}

func NewDetectPestImageController(detectPestImageService services.DetectPestImageServiceInterface) DetectPestImageControllerInterface {
	return &DetectPestImageController{DetectPestImageService: detectPestImageService}
}

func (dpc *DetectPestImageController) DetectPestImage(c *gin.Context) {

	ctx := context.Background()

	responseApiPython, err := dpc.DetectPestImageService.DetectPestImage(c.Request)

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
