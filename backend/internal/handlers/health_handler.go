package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishnu777000/OpsLens/backend/internal/services"
)

func HealthHandler(c *gin.Context) {
	response := services.GetHealth()
	c.JSON(http.StatusOK, response)
}
