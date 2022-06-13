package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"telecommunication/controllers"
	"telecommunication/entities"
	"telecommunication/services"

	"github.com/gin-gonic/gin"
)

var (
	smsService    services.SMSService       = services.NewSMSService()
	smsController controllers.SMSController = controllers.NewSMSController(smsService)
)

func SMS(c *gin.Context) {
	var (
		res        = entities.SMSResponse{}
		serviceRes string
	)
	serviceRes = smsController.SendSMS(c)
	if serviceRes != "" {
		var objmap map[string]*json.RawMessage
		if err := json.Unmarshal([]byte(serviceRes), &objmap); err != nil {
			log.Fatal(err)
		}
		res.Data = objmap
	}
	res.Message = "success"
	res.Success = true
	c.JSON(http.StatusOK, res)
}
