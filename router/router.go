package router

import "github.com/gin-gonic/gin"

func StartRouter() {

	router := gin.Default()

	router.Run("localhost:8080")
}
