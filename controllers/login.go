package controllers

import (
	"telecommunication/entities"
	"telecommunication/services"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) (jwtToken string)
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

func NewLoginController(loginService services.LoginService, jwtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) (jwtToken string) {
	var credentials entities.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return
	}
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)

	if isAuthenticated {
		jwtToken = c.jwtService.GenerateToken(credentials.Username, true)
	}
	return
}
