package main

import (
	"io"
	"os"

	"github.com/danxcraig/demob-main/controllers"
	"github.com/danxcraig/demob-main/middlewares"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	homepageRoute := server.Group("/")
	{
		homepageRoute.GET("/", controllers.Get)
	}

	testRoutes := server.Group("/test")
	{
		testRoutes.GET("/", controllers.Test)
	}

	server.Run(":9090")
}
