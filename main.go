package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/wowBlaBla/FizzBuzz-API/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/fizzbuzz", controller.FizzBuzzController)

	router.Run(":3001")
}
