package controller

import (
	"net/http"
	"rest-go/models"

	"github.com/gin-gonic/gin"
)

var fitness = []models.Fitness{
	{Status: "healthy"},
}

func GetFitness(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fitness)
}
