package handlers

import (
	"fmt"
	"net/http"
	"telecommunication/controllers"
	"telecommunication/services"

	"github.com/gin-gonic/gin"
)

var (
	dummyService    services.DummyService       = services.New()
	dummyController controllers.DummyController = controllers.New(dummyService)
)

func Dummy(c *gin.Context) {
	res, err := dummyController.Dummy(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, res)
}
