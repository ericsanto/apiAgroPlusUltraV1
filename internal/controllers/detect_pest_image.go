package controllers

import (
	"context"
	"errors"
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
		if errors.Is(err, ctx.Err()) {
			myerror.HttpErrors(http.StatusRequestTimeout, err.Error(), c)
			return
		}

		if errors.As(err, &sarama.ConsumerError{}) {
			myerror.HttpErrors(http.StatusBadGateway, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, responseApiPython)
}
