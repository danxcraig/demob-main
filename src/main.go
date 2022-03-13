package main

import (
	"github.com/danxcraig/demob-main/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.Get)

	r.Run(":9090")
}
