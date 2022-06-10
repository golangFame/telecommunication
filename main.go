package main

import (
	"io"
	"net/http"
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

	routesWithBasicAuth := router.Group("/basicauth")
	routesWithBasicAuth.Use(middlewares.BasicAuth())

	// router.Use(gin.Logger())

	router.GET("/", handlers.Home)
	router.GET("/dummy", handlers.Dummy)
	router.POST("/login", handlers.Login)

	routesWithBasicAuth.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	router.Run(":8080")
}
