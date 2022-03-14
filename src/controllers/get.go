package controllers

import (
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "server is ok",
	})
}

func Test(ctx *gin.Context) {
	var names []string = []string{"D", "P", "E"}
	data := gin.H{
		"titles": "the main title for the page",
		"names":  names,
	}
	ctx.HTML(200, "index.html", data)
}
