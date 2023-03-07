package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FizzBuzzController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
