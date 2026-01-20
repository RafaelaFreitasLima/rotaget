package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	router.GET("/fitness", controller.GetFitness)
	router.Run("localhost:8080")
}
