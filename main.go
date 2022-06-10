package main

import (
	"io"
	"os"
	"telecommunication/handlers"
	"telecommunication/middlewares"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	router := gin.New()

	router.Use(gin.Recovery(), middlewares.Logger(),
		gindump.Dump())

	// router.Use(gin.Logger())

	router.GET("/", handlers.Home)
	router.GET("/dummy", handlers.Dummy)
	router.POST("/login", handlers.Login)

	router.Run(":8080")
}
