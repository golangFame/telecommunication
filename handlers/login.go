package handlers

import (
	"net/http"
	"telecommunication/controllers"
	"telecommunication/services"

	"github.com/gin-gonic/gin"
)

var (
	loginService    services.LoginService       = services.NewLoginService()
	jwtService      services.JWTService         = services.NewJWTService()
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func Login(c *gin.Context) {

	token := loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}
