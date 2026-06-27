package services

import "github.com/gin-gonic/gin"

func GetHealth() gin.H {
	return gin.H{
		"status":  "healthy",
		"service": "OpsLens Backend",
		"version": "0.1.0",
	}
}
