package route

import (
	"github.com/gin-gonic/gin"
	controller "github.com/wowBlaBla/FizzBuzz-API/controllers"
)

func InitFizzBuzzRouter(r *gin.Engine) {
	r.POST("/fizzbuzz", controller.FizzBuzzController)
}
