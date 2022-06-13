package controllers

import (
	"net/http"
	"telecommunication/entities"
	"telecommunication/services"

	"github.com/gin-gonic/gin"
)

type SMSController interface {
	SendSMS(ctx *gin.Context) string
}

type smsController struct {
	smsService services.SMSService
}

func NewSMSController(smsService services.SMSService) SMSController {
	return &smsController{
		smsService: smsService,
	}
}

func (c *smsController) SendSMS(ctx *gin.Context) (response string) {
	var err error
	serviceName := ctx.GetHeader("X-Service-Name")
	serviceVersion := ctx.GetHeader("X-Service-Version")

	_ = ctx.GetHeader("Origin") // Do something with service

	switch serviceName {
	case "KALERYA":
		switch serviceVersion {
		case "4.0.0":
			data := entities.KaleyraSMSRequest{}
			if err = ctx.ShouldBind(&data); err != nil {

				ctx.JSON(http.StatusUnprocessableEntity, err)
				return
			}
			response = c.smsService.SendKaleryaSMS(data)
		}
	}
	return
}
