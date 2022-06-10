package controllers

import (
	"telecommunication/entities"
	"telecommunication/services"
	"telecommunication/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DummyController interface {
	Dummy(ctx *gin.Context) (res entities.Dummy, err error)
}

type controller struct {
	service services.DummyService
}

var validate *validator.Validate

func New(s services.DummyService) DummyController {
	validate = validator.New()
	validate.RegisterValidation("dummy", validators.ValidateDummyTag)
	return &controller{
		service: s,
	}
}

func (c *controller) Dummy(ctx *gin.Context) (res entities.Dummy, err error) {
	var (
		req = entities.Dummy{}
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	err = validate.Struct(req)
	if err != nil {
		return
	}
	res = c.service.Dummy()
	return
}
