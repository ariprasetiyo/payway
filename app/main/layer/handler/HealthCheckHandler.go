package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ok",
		"data":    "application running",
	})
}
