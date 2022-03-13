package main

import (
	"github.com/danxcraig/demob-main/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", controllers.Get)

	server.Run(":9090")
}
