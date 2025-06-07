package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
)

func DetectPestImageController(c *gin.Context) {

	ctx := context.Background()

	responseApiPython, err := services.DetectPestImage(c.Request)

	if err != nil {
		switch {
		case errors.Is(err, ctx.Err()):
			myerror.HttpErrors(http.StatusRequestTimeout, err.Error(), c)

		case errors.As(err, &sarama.ConsumerError{}):
			log.Println(err.Error())
			myerror.HttpErrors(http.StatusBadGateway, "erro ao consultar servidor externo", c)

		case errors.Is(err, myerror.ErrImageSizeToLarge):
			myerror.HttpErrors(413, err.Error(), c)

		case errors.Is(err, myerror.ErrUnsupportedImageType):
			myerror.HttpErrors(http.StatusUnsupportedMediaType, err.Error(), c)

		default:
			myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		}
	}

	c.JSON(http.StatusOK, responseApiPython)
}
