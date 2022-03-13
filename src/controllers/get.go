package controllers

import (
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "server is ok",
	})
}
