package controllers

import (
	"money-game-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activities
	err := models.GetAllActivities(&activities)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, activities)
	}
}
